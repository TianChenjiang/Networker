package web

import (
	"citicup-admin/internal/web/e"
	"citicup-admin/schema"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)


func GetUserList(c *gin.Context) {
	var (
		appG = Gin{C: c}
	)
	list, err := serv.GetAllUser(*c)
	if err != nil {
		log.Print("err")
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USERLIST, nil)
		return
	}
	appG.OK(list)
}


func GetUserByID(c *gin.Context) {
	var (
		appG  = Gin{C: c}
		id, _ = strconv.Atoi(c.Param("id"))
	)

	user, err := serv.GetUserById(*c, uint(id))
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USER, nil)
		return
	}
	appG.OK(user)
}

func Register(c *gin.Context) {
	var (
		appG = Gin{C: c}
		schema  schema.User
	)
	err := c.Bind(&schema)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_CREATE_USER, nil)
		return
	}

	u, err := serv.CreateUser(*c, &schema)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_USER, nil)
		return
	}

	appG.OK(u)
}

func EditUser(c *gin.Context) {
	var (
		appG = Gin{C: c}
		schema  schema.User
	)
	err := c.Bind(&schema)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_EDIT_USER, nil)
		return
	}

	u, err := serv.UpdateUser(*c, &schema)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_USER, nil)
		return
	}

	appG.OK(u)
}

func DeleteUser(c *gin.Context) {
	var (
		appG = Gin{C: c}
		id, _ = strconv.Atoi(c.Param("id"))
	)

	user, err := serv.DeleteUser(*c, uint(id))
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_USER, nil)
		return
	}
	appG.OK(user)
}








