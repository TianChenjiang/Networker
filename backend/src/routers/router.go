package routers

import (
	"github.com/gin-gonic/gin"
	"networker/backend/src/pkg/setting"

	"go-web-demo/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	//获取token
	r.GET("/auth", api.GetAuth)


	//apiv1 := r.Group("/api/v1")


	return r
}
