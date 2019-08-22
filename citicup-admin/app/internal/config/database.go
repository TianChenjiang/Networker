package config

import (
	"citicup-admin/internal/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

//数据库连接池
type DBLink struct {
	*gorm.DB
}

//建立数据库连接池
func SetUpDBLink(dbConfig *DataBase) *DBLink {
	var err error
	//数据库连接
	db, err = gorm.Open(dbConfig.Type,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.dbName,
		))
	if err != nil {
		log.Fatalf("database err: %v", err)
	}
	db.SingularTable(true)
	//自动映射
	db.AutoMigrate(&model.User{}, &model.Company{})
	log.Print("auto complete the databases")
	return &DBLink{db}
}
