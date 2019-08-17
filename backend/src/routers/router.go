package routers

import (
	"github.com/gin-gonic/gin"
	"networker/backend/src/routers/api/controller"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	user := controller.User{} //todo


	v1 := r.Group("/api/v1")

	{
		v1.POST("/users", user.CreateUser) //TODO
	}

	return r
}


