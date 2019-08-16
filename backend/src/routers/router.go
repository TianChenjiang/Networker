package routers

import (
	"github.com/gin-gonic/gin"
	"src/pkg/setting"

	//"go-web-demo/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	//获取token
	//TODO: 这里的不太对
	//r.GET("/auth", api.GetAuth)

	//apiv1 := r.Group("/api/v1")

	return r
}
