package gormplus


import (
	"fmt"
	"log"
	"networker/backend/src/pkg/setting"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

var db *gorm.DB

//// New 创建DB实例
//func New(c *Config) (*DB, error) {
//	db, err := gorm.Open(c.DBType, c.DSN)
//	if err != nil {
//		return nil, err
//	}
//
//	if c.Debug {
//		db = db.Debug()
//	}
//
//	err = db.DB().Ping()
//	if err != nil {
//		return nil, err
//	}
//
//	db.DB().SetMaxIdleConns(c.MaxIdleConns)
//	db.DB().SetMaxOpenConns(c.MaxOpenConns)
//	db.DB().SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)
//	return &DB{db}, nil
//}


// Wrap 包装gorm
func Wrap(db *gorm.DB) *DB {
	return &DB{db}
}

// DB gorm扩展DB
type DB struct {
	*gorm.DB
}


func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	//回调 自动更新更新，删除时间
	db.Callback().Create().Replace("gormplus:update_time_stamp", UpdateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gormplus:update_time_stamp", UpdateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gormplus:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

//得到当前db
func GetDB() *gorm.DB {
	return db
}


func CloseDB() {
	defer db.Close()
}

func UpdateTimeStampForCreateCallback(scope *gorm.Scope) {
	if ! scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}

}

func UpdateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gormplus:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}


func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gormplus:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

