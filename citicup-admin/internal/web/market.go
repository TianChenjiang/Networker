package web

import (
	"citicup-admin/internal/web/e"
	"citicup-admin/schema"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllMarketCondition(c *gin.Context)  {
	var (
		appG = Gin{C: c}
	)

	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	PageNum, _  := strconv.Atoi(c.Query("pageNum"))


	res, err := serv.GetAllMarket(PageNum, PageSize)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INTERNAL_ERROR, nil)
	}

	appG.OK(res)
	return
}

func GetConcernedMarketCondition(c *gin.Context)  {
	var (
		appG = Gin{C: c}
	)

	//获得当前用户token
	user, code, err:= serv.GetUserByToken(*c)
	if err != nil || code != e.SUCCESS{
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}

	//获得该用户所有关注公司
	companyList,err := serv.GetConcerned(0, 10000, user.ID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INTERNAL_ERROR, nil)
	}

	marketList := make([]schema.Market, 0)

	for i := 0; i < len(companyList); i++{
		market, _ := serv.GetMarket(companyList[i].ID)
		marketList = append(marketList, market.Model2Schema())
	}

	appG.OK(marketList)
}

func GetMarketConditionBySymbol(c *gin.Context)  {
	var (
		appG = Gin{C: c}
		symbol = c.Query("symbol")
	)

	//获得当前用户token
	user, code, err:= serv.GetUserByToken(*c)
	if err != nil || code != e.SUCCESS {
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}

	market, isConcerned, err := serv.GetMarketConditionBySymbol(user.ID, symbol)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INTERNAL_ERROR, nil)
	}

	data := make(map[string]interface{})
	data["isConcerned"] = isConcerned
	data["market"] = market.Model2Schema()

	appG.OK(data)
	
}

func InsertMarket(c *gin.Context)  {
	var (
		appG = Gin{C: c}
		schema  schema.InsertMarketParam
	)

	err := c.BindJSON(&schema)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.BAD_REQUEST, nil)
	}

	market, err := serv.InsertMarket(schema)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INTERNAL_ERROR, nil)
	}



	appG.OK(market.Model2Schema())

}

