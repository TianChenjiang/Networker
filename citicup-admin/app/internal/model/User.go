package model

//用户实体
type User struct {
	ID           uint     `gorm:"column:id;AUTO_INCREMENT"`
	UserName     string   `gorm:"column:username"`  //todo index?
	Password     string   `gorm:"column:password"`
	Email        string   `gorm:"column:email"`
	Phone        string   `gorm:"column:phone"`
}

func (u User) TableName() string {
	return "user"
}
