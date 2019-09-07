package model

import (
	"citicup-admin/schema"
)

//用户实体
type User struct {
	ID       uint   `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Email    string `gorm:"column:email"`
	Phone    string `gorm:"column:phone"`
	Avatar   string `gorm:"column:avatar"`
	Companies []Company `gorm:"many2many:user_company;"`
}



func (u *User) TableName() string {
	return "user"
}

func (u *User) Model2Schema() (result schema.User) {
	result = schema.User{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
		Phone:    u.Phone,
		Email:    u.Email,
	}
	return
}

type Users []User

func (u Users) Model2Schema() (result []schema.User) {
	result = make([]schema.User, len(u))
	for i, item := range u {
		result[i] = item.Model2Schema()
	}
	return
}
