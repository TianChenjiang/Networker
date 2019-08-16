package controller

import (
	"github.com/gin-gonic/gin"
	"networker/backend/src/bl"
	"networker/backend/src/ginplus"
	"networker/backend/src/schema"
)

type user bl.IUser

func NewUser(user bl.IUser) *User {
	return &User{
		UserBl:  user,
	}
}

type User struct {
	UserBl	bl.IUser
}

func (user *User) CreateUser (c *gin.Context)  {
	var item schema.User
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	nitem, err := user.UserBl.Create(ginplus.NewContext(c), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nitem) //todo
}


