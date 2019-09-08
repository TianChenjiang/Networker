package main

import "github.com/jinzhu/gorm"

// User 包含多个 emails, UserID 为外键
type User struct {
	gorm.Model
	Emails []Email
}

type Email struct {
	gorm.Model
	Email  string
	UserID uint
}
type db struct {
	*gorm.DB
}

func main() {
}
