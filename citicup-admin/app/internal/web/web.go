package web

import (
	_ "citicup-admin/docs" // docs is generated by Swag CLI, you have to import it.
	"citicup-admin/internal/config"
	"citicup-admin/internal/service"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"strconv"
	"strings"
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
		Addr:    strings.Join([]string{":", strconv.Itoa(c.App.Port)}, ""),
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
	e.StaticFS("/upload/image", http.Dir("library/pic/"))
	
	g := e.Group("/api")
	{
		userRouter := g.Group("/users")
		//userRouter.Use(middleware.JWT())
		{
			userRouter.GET("", GetUserList)
			userRouter.GET("/id/:id", GetUserByID)
			userRouter.POST("", Register)
			userRouter.PUT("", EditUser)
			userRouter.DELETE("/:id", DeleteUser)
			userRouter.POST("/login", UserLogin)
			userRouter.POST("/change/password", ChangePassword)
			userRouter.POST("/upload", UploadAvatar)
		}

		companyRouter := g.Group("/companies")
		{
			companyRouter.GET("", GetCompanies)
			companyRouter.GET("/:id", GetCompanyById)
			companyRouter.POST("", NewCompany)
			companyRouter.DELETE("/:id", DeleteCompanyById)
		}
	}
}
