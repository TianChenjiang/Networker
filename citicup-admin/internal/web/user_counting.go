package web

import (
	"citicup-admin/internal/web/e"
	"github.com/gin-gonic/gin"
	"net/http"
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
		//TODO: 完成不同角色的用户数量查询
	} else {
		//bad request
		appG.Response(http.StatusBadRequest, e.BAD_REQUEST, nil)
		return
	}
	appG.OK(nil)

}

