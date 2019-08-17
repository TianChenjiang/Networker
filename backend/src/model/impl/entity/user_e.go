package entity

import "networker/backend/src/schema"

//用户实体
type User struct {
	ID           uint     `gorm:"column:id;AUTO_INCREMENT"`
	UserName     string   `gorm:"column:username"`  //todo index?
	Password     string   `gorm:"column:password"`
	Role         string   `gorm:"column:role"` //user admin
	Email        string   `gorm:"column:email"`
	Phone        string   `gorm:"column:phone"`
}


func (u User) TableName() string {
	return "user"
}

// SchemaUser 用户对象
type SchemaUser schema.User

// ToUser 转换为用户实体
func (s SchemaUser) ToUser() User {
	item := User{
		ID:       s.ID,
		UserName: s.Username,
		Password: s.Password,
		Role:     s.Role,
		Email:    s.Email,
		Phone:    s.Phone,
	}

	return item
}

//用户实体转为用户对象
func (u User) ToSchemaUser() schema.User  {
	item := schema.User{
		ID:       u.ID,
		Username: u.UserName,
		Password: u.Password,
		Role:     u.Role,
		Phone:    u.Phone,
		Email:    u.Email,
	}
	return item
}

type Users []User
//转换为用户对象列表
func (u Users) ToSchemaUsers() []schema.User {
	list := make([]schema.User, len(u))
	for i, item := range u {
		list[i] = item.ToSchemaUser()
	}
	return list
}















