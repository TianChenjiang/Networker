package web

import (
	"citicup-admin/internal/config"
	"citicup-admin/internal/service"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

var (
	serv *service.Service
)

func New(c *config.Config, s *service.Service) (httpSrv *http.Server) {
	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	routeSetUp(engine)

	httpSrv = &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
	serv = s
	go func() {
		// service connections
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	return
}

func routeSetUp(e *gin.Engine) {
	//swagger
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g := e.Group("/api")
	{
		u := g.Group("/users")
		{
			u.GET("", GetUserList)
		}
	}
}
