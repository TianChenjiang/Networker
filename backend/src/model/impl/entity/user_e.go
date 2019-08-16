package entity

import "networker/backend/src/schema"

//用户实体
type User struct {
	ID           uint     `gormplus:"column:id;AUTO_INCREMENT"`
	UserName     string   `gormplus:"column:username"`  //todo index?
	Password     string   `gormplus:"column:password"`
	Role         string   `gormplus:"column:role"` //user admin
	Email        string   `gormplus:"column:email"`
	Phone        string   `gormplus:"column:phone"`
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















