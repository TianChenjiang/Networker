package web

import (
	_ "citicup-admin/docs" // docs is generated by Swag CLI, you have to import it.
	"citicup-admin/internal/config"
	"citicup-admin/internal/middleware"
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
	e.StaticFS("/image", http.Dir("library/pic/")) //前端访问已上传的图片

	g := e.Group("/api")
	{
		userRouter := g.Group("/users")
		//userRouter.Use(middleware.JWT())
		userRouter.Use(middleware.CorsMiddleware())
		{
			userRouter.GET("", GetUserList)
			userRouter.GET("/id/:id", GetUserByID)
			userRouter.POST("/register", Register)
			userRouter.PUT("", EditUser)
			userRouter.DELETE("/:id", DeleteUser)
			userRouter.POST("/login", UserLogin)
			userRouter.PUT("/password", ChangePassword)
			userRouter.POST("/upload", UploadAvatar)
			userRouter.GET("/mark", MarkAsConcerned)
			userRouter.GET("/unmark", CancelMarkAsConcerned)
			userRouter.GET("/concerned", GetConcerned)
			userRouter.GET("/concerned/market", GetConcernedMarketCondition)

		}

		companyRouter := g.Group("/companies")
		{
			companyRouter.GET("", GetCompanies)
			companyRouter.GET("/:id", GetCompanyById)
			companyRouter.POST("", NewCompany)
			companyRouter.PUT("", UpdateCompany)
			companyRouter.DELETE("/:id", DeleteCompanyById)
		}

		investorRouter := g.Group("/investor")
		{
			investorRouter.POST("/register", InvestorRegister)
			investorRouter.POST("/upload", UploadRegisterInfo)
			investorRouter.POST("/login", InvestorLogin)
		}

		marketRouter := g.Group("/market")
		{
			marketRouter.GET("/all", GetAllMarketCondition)
			marketRouter.POST("/insert", InsertMarket)
			marketRouter.GET("/symbol", GetMarketConditionBySymbol)
		}
	}
}
