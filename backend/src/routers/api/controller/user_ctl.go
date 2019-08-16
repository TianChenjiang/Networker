package controller

import (
	"github.com/gin-gonic/gin"
	"networker/backend/src/bl"
	"networker/backend/src/ginplus"
	"networker/backend/src/schema"
)


func NewUser(user bl.IUser) *User {
	return &User{
		UserBl:  user,
	}
}

type User struct {
	UserBl	bl.IUser
}

func (u *User) Create (c *gin.Context)  {
	var item schema.User
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	nitem, err := u.UserBl.Create(ginplus.NewContext(c), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nitem) //todo
}

func (a *User) Update(c *gin.Context) {
	var item schema.User
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	nitem, err := a.UserBll.Update(ginplus.NewContext(c), c.Param("id"), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nitem.CleanSecure())
}


