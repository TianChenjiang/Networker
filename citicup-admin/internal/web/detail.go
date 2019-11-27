package web

import (
	"citicup-admin/schema"
	"github.com/gin-gonic/gin"
)

func GetCandleBySymbol(c *gin.Context) {
	var (
		appG   = Gin{C: c}
		symbol = c.Query("symbol")
	)
	market, _, _, _ := serv.GetMarketConditionBySymbol(1, symbol)
	candle := schema.Candle{
		Close:   market.Close,
		Highest: market.Highest,
		Lowest:  market.Lowest,
		Open:    market.Open,
		Vol:     market.Vol,
	}
	appG.OK(candle)
}


func GetFinanceBySymbol(c *gin.Context) {
	var (
		appG   = Gin{C: c}
		symbol = c.Query("symbol")
	)
	finance, _ := serv.GetFinanceBySymbol(symbol)
	appG.OK(finance.Model2Schema())
}


func GetPledgeBySymbol(c *gin.Context) {
	var (
		appG   = Gin{C: c}
		symbol = c.Query("symbol")
	)
	pledge, _ := serv.GetPledgeBySymbol(symbol)
	appG.OK(pledge.Model2Schema())
}
