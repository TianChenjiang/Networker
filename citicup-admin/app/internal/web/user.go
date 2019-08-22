package web

import (
	"github.com/gin-gonic/gin"
	"log"
)

// @Tags User
// @Summary 获取所有用户
// @Description Get All Users
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Users
// @Failure 404 {string} string "No Users"
// @Failure 500 {string} string "Internal Error"
// @Router /api/users [get]
func GetUserList(c *gin.Context) {
	var (
		appG = Gin{C: c}
	)
	list, err := serv.GetAllUser(*c)
	if err != nil {
		log.Print("err")
	}
	appG.OK(list)
}
