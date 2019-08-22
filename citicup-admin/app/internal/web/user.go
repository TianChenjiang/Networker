package web

import (
	"citicup-admin/internal/web/e"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USERLIST, nil)
		return
	}
	appG.OK(list)
}


func GetUserByID(c *gin.Context) {
	var (
		appG = Gin{C: c}
	)



}








