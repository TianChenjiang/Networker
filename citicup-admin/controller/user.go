package controller

import (
	"citicup-admin/common"
	"citicup-admin/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	common.Controller
}

var userService service.UserService

//路由注册
func (ctrl *UserController) Router(router *gin.Engine) {

	r := router.Group("user")
	{

		r.GET("", ctrl.findOne)
	}

}

func (ctrl *UserController) findOne(ctx *gin.Context) {

	userId := ctx.Query("id")

	ret := userService.FindOne(userId)
	common.ResultOk(ctx, ret)
}
