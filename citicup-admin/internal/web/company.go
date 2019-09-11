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
// @Param pageSize query int true "页大小"
// @Param pageNum query int true "页号"
// @Success 200 {object} model.Companies
// @Failure 404 {string} string "Resource not found"
// @Failure 500 {string} string "Internal Error"
// @Router /api/companies/list [get]
func GetCompanies(c *gin.Context) {
	var (
		appG = Gin{C: c}
	)

	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	list, err := serv.GetAllCompaniesPaging(c, pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INTERNAL_ERROR, nil)
		return
	}
	appG.OK(list)
}

// @Tags Company
// @Summary 模糊搜索公司
// @Accept  json
// @Produce  json
// @Param key query string true "关键字，按照关键字搜索名称"
// @Param pageSize query int true "页大小"
// @Param pageNum query int true "页号"
// @Success 200 {object} model.Company
// @Failure 404 {string} string "Resource not found"
// @Failure 500 {string} string "Internal Error"
// @Router /api/companies/query [get]
func QueryCompanies(c *gin.Context) {
	var (
		appG = Gin{C: c}
	)

	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	key := c.Query("key")
	list, err := serv.QueryCompanies(c, key, pageNum, pageSize)
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
// @Param id query string true "公司id"
// @Success 200 {object} model.Company
// @Failure 404 {string} string "Resource not found"
// @Failure 500 {string} string "Internal Error"
// @Router /api/companies [get]
func GetCompanyById(c *gin.Context) {
	var (
		appG = Gin{C: c}
		id   = c.Query("id") //Path  :id
	)
	i, _ := strconv.Atoi(id)
	company, err := serv.GetCompanyById(c, uint(i))
	if company.ID <= 0 {
		appG.Response(http.StatusNotFound, e.RESOURCE_NOT_FOUND, nil)
		return
	}
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INTERNAL_ERROR, nil)
		return
	}
	appG.OK(company)
}

// @Tags Company
// @Summary 更新公司信息
// @Description Update the company information
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 404 {string} string "Resource not found"
// @Failure 500 {string} string "Internal Error"
// @Router /api/companies [PUT]
func UpdateCompany(c *gin.Context) {
	var (
		appG   = Gin{C: c}
		scheme schema.Company
	)
	err := c.BindJSON(&scheme)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.BAD_REQUEST, nil)
		return
	}
	err_ := serv.UpdateCompany(c, &scheme)
	if err_ != nil {
		appG.Response(http.StatusNotFound, e.BAD_REQUEST, nil)
	}
	appG.OK(nil)
}

// @Tags Company
// @Summary 删除指定公司信息
// @Description Delete company
// @Param id query string true "公司id"
// @Accept  json
// @Success 200
// @Failure 404 {string} string "Resource not found"
// @Failure 500 {string} string "Internal Error"
// @Router /api/companies [DELETE]
func DeleteCompanyById(c *gin.Context) {
	var (
		appG = Gin{C: c}
		id   = c.Query("id") //Path  :id
	)
	i, _ := strconv.Atoi(id)
	err := serv.DeleteCompany(c, uint(i))
	if err != nil {
		appG.Response(http.StatusNotFound, e.INTERNAL_ERROR, nil)
		return
	}
	appG.OK(nil)
}

// @Tags Company
// @Summary 添加新的公司
// @Description Delete company
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Company
// @Failure 404 {string} string "Resource not found"
// @Failure 500 {string} string "Internal Error"
// @Router /api/companies [POST]
func NewCompany(c *gin.Context) {
	var (
		appG   = Gin{C: c}
		scheme schema.Company
	)
	err := c.BindJSON(&scheme)
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
