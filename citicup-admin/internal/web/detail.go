package web

import "github.com/gin-gonic/gin"

func GetFinanceBySymbol(c *gin.Context)  {
	var (
		appG = Gin{C: c}
		symbol = c.Query("symbol")
	)

	appG.OK(symbol)

}
