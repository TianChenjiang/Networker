package web

import (
	"citicup-admin/internal/web/e"
	"citicup-admin/schema"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Tags Company
// @Summary 获取所有公司
// @Description Get All Companies.
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Companies
// @Failure 404 {string} string "Resource not found"
// @Failure 500 {string} string "Internal Error"
// @Router /api/company [get]
func GetCompanies(c *gin.Context) {
	var (
		appG = Gin{C: c}
	)
	list, err := serv.GetAllCompanies(c)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INTERNAL_ERROR, nil)
		return
	}
	appG.OK(list)
}

// @Tags Company
// @Summary 根据Id获取公司信息
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Companies
// @Failure 404 {string} string "Resource not found"
// @Failure 500 {string} string "Internal Error"
// @Router /api/company/{id} [get]
func GetCompanyById(c *gin.Context) {
	var (
		appG = Gin{C: c}
		id   = c.Param("id") //Path  :id
	)
	i, _ := strconv.Atoi(id)
	company, err := serv.GetCompanyById(c, uint(i))
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INTERNAL_ERROR, nil)
		return
	}
	appG.OK(company)
}

func UpdateCompany(c *gin.Context) {
	//var(
	//	appG = Gin{C:c},
	//
	//)
}

func DeleteCompanyById(c *gin.Context) {
	var (
		appG = Gin{C: c}
		id   = c.Param("id") //Path  :id
	)
	i, _ := strconv.Atoi(id)
	err := serv.DeleteCompany(c, uint(i))
	if err != nil {
		appG.Response(http.StatusNotFound, e.INTERNAL_ERROR, nil)
		return
	}
	appG.OK(nil)
}

func NewCompany(c *gin.Context) {
	var (
		appG   = Gin{C: c}
		scheme schema.Company
	)
	err := c.Bind(&scheme)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.BAD_REQUEST, nil)
		return
	}
	newCom, err_ := serv.NewCompany(c, &scheme)
	if err_ != nil {
		appG.Response(http.StatusNotFound, e.BAD_REQUEST, nil)
	}
	appG.OK(newCom)
}
