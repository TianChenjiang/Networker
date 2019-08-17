package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"networker/backend/src/app"
	"networker/backend/src/bl"
	"networker/backend/src/ginplus"
	"networker/backend/src/pkg/e"
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

func (u *User) CreateUser (c *gin.Context)  {
	var (
		appG = app.Gin{C: c}
		form schema.CreateUserForm
		err error
	)


	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
	}

	item := schema.User{
		ID:       form.ID,
		Username: form.Username,
		Password: form.Password,
		Role:     form.Role,
		Phone:    form.Phone,
		Email:    form.Email,
	}


	fmt.Println(item)
	fmt.Println(u)

	_, err = u.UserBl.Create(ginplus.NewContext(c), item)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}


