package web

import (
	"citicup-admin/internal/model"
	"citicup-admin/internal/web/e"
	"citicup-admin/schema"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllMarketCondition(c *gin.Context)  {
	var (
		appG = Gin{C: c}
		PaginationScheme schema.Pagination
	)

	err := c.BindJSON(&PaginationScheme)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_CREATE_USER, nil)
		return
	}

	res, err := serv.GetAllMarket(PaginationScheme.PageNum, PaginationScheme.PageSize)
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
	if err != nil {
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}

	//获得该用户所有关注公司
	companyList,err := serv.GetConcerned(user.ID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INTERNAL_ERROR, nil)
	}

	marketList := make([]model.Market, 0)

	for i := 0; i < len(companyList); i++{
		market, _ := serv.GetMarket(companyList[i].Market.ID)
		marketList = append(marketList, market)
	}

	appG.OK(marketList)
}

func GetMarketConditionBySymbol(c *gin.Context)  {
	
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

	appG.OK(market)

}

