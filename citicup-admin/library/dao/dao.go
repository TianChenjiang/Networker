package dao

//封装gorm,使得CURD更加简洁
import "github.com/jinzhu/gorm"

func FindAll(db *gorm.DB, model interface{}, out interface{}) {
	db.Model(model).Find(out)
}

func Save(db *gorm.DB, model interface{}, entity interface{}) {
	db.Model(model).Save(entity)
}

func Update(db *gorm.DB, model interface{}, entity interface{}) {
	db.Model(model).Update(entity)
}
