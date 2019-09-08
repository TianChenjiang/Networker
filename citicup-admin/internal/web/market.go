package web

import (
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
	
}

func GetMarketConditionBySymbol(c *gin.Context)  {
	
}

