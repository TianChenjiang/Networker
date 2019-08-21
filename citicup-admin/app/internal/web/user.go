package web

import (
	"github.com/gin-gonic/gin"
	"log"
)

// @Summary 获取用户列表
// @Produce json
// @Success 200 {object} render.JSON
// @Router /api/Users [get]
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
