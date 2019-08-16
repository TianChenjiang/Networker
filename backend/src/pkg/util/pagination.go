package util

import (
	"github.com/gin-gonic/gin"
	//"github.com/go-ini/ini"
	"github.com/Unknwon/com"
	"networker/backend/src/pkg/setting"
)

//分页页码
func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}
	return result
}