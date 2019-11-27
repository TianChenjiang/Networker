package web

import (
	_ "citicup-admin/docs" // docs is generated by Swag CLI, you have to import it.
	"citicup-admin/internal/config"
	//"citicup-admin/internal/middleware"
	"citicup-admin/internal/service"
	"github.com/gin-contrib/cors"
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

	allowedHeaders := []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"}
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = allowedHeaders
	engine.Use(cors.New(config))
	engine.Use(cors.Default())

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
		//userRouter.Use(middleware.CorsMiddleware())
		{
			userRouter.GET("", GetUserList)
			userRouter.GET("/count",FetchUser)
			userRouter.POST("/check",CheckInvestorReq)
			userRouter.POST("/investor",FindAllInvestorsByStatus)
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
			companyRouter.GET("/list", GetCompanies)
			companyRouter.GET("/query", QueryCompanies)
			companyRouter.GET("/query/brief", BriefQueryCompanies)
			companyRouter.GET("", GetCompanyById)
			companyRouter.POST("", NewCompany)
			companyRouter.PUT("", UpdateCompany)
			companyRouter.DELETE("", DeleteCompanyById)

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

		detailRouter := g.Group("/detail")
		{
			detailRouter.GET("/finance", GetFinanceBySymbol)
			detailRouter.GET("/pledge", GetPledgeBySymbol)
			detailRouter.GET("/candle", GetCandleBySymbol)
		}
	}
}
