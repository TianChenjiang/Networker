package web

import (
	"citicup-admin/internal/web/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Tags User
// @Summary 查询用户数量
// @Accept  json
// @Produce  json
// @Param role query string true "根据该query分辨普通用户和质押方.用USER表示普通用户,用INVESTOR表示质押方.其他输入无效"
// @Success 200 {int} int "用户数量"
// @Failure 400 {string} string "错误的query"
// @Failure 500 {string} string "Internal Error"
// @Router /api/users/count [get]
func FetchUser(c *gin.Context) {
	var (
		appG = Gin{C: c}
		role = c.Query("role")
	)
	if role == "USER" || role == "INVESTOR" {
		result, _ := serv.FetchUserCounting(c, role)
		appG.OK(&result)
		return
	} else {
		//bad request
		appG.Response(http.StatusBadRequest, e.BAD_REQUEST, nil)
		return
	}
}

// @Tags User
// @Summary 确认质押方的注册申请结果
// @Accept  json
// @Produce  json
// @Param checkResult query bool true "true 同意; false: 不同意，那么该用户将被删除"
// @Param id query int true "用户id"
// @Success 200 {string} string
// @Failure 400 {string} string "错误的query"
// @Failure 500 {string} string "Internal Error"
// @Router /api/users/check [POST]
func CheckInvestorReq(c *gin.Context) {
	var (
		appG = Gin{C: c}
		stat = c.Query("checkResult")
		id   = c.Query("id") //query  :id
	)
	i, _ := strconv.Atoi(id)
	s, _ := strconv.ParseBool(stat)
	serv.CheckInvestorReq(c, uint(i), s)
	appG.OK(nil)
}

// @Tags User
// @Summary 获取质押方列表,可以根据是否注册成功进行筛选
// @Accept  json
// @Produce  json
// @Param status query bool true "true 已审批通过; false待审批"
// @Param pageNum query int true "页号"
// @Param pageSize query int true "页大小"
// @Success 200 {object} model.Investors
// @Failure 400 {string} string "错误的query"
// @Failure 500 {string} string "Internal Error"
// @Router /api/users/investor [POST]
func FindAllInvestorsByStatus(c *gin.Context) {
	var (
		appG = Gin{C: c}
		stat = c.Query("status")
	)
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	s, _ := strconv.ParseBool(stat)
	list, err := serv.FetchInvestorListByStatus(c, s, pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.BAD_REQUEST, nil)
		return
	}
	appG.OK(list)
}
