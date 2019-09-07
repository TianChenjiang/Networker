package web

import (
	"citicup-admin/internal/web/e"
	"citicup-admin/schema"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InvestorRegister(c *gin.Context)  {
	var (
		appG = Gin{C: c}
		schema  schema.Investor
	)
	err := c.BindJSON(&schema)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.BAD_REQUEST, nil)
		return
	}

	i, err := serv.CreateInvestor(&schema)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INTERNAL_ERROR, nil)
		return
	}

	appG.OK(i)
}
