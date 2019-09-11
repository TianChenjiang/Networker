package service

import "github.com/gin-gonic/gin"

//账户管理

func (s *Service) FetchUserCounting(c *gin.Context, role string) (result int, err error) {
	if role == "USER" {
		result, err = s.dao.CountUsers()
	} else {
		result, err = s.dao.CountInvestor()
	}
	return
}
