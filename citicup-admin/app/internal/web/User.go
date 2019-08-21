package web

import (
	"github.com/CaribouW/gin-common-lib/library/ecode"
	"github.com/CaribouW/gin-common-lib/library/render"
	"github.com/gin-gonic/gin"
	"log"
)

// @Summary 用户列表
// @Produce json
// @Success 200 {object} render.JSON
// @Router /api/Users [get]
func GetUserList(c *gin.Context) {
	r := render.New(c)
	list, err := serv.GetAllUser(*c)
	if err != nil {
		log.Print("err")
	}
	r.JSON(list, ecode.OK)
}
