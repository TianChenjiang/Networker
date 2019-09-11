package service

import (
	"citicup-admin/schema"
	"github.com/gin-gonic/gin"
)

//账户管理

//获取用户数量
func (s *Service) FetchUserCounting(c *gin.Context, role string) (result int, err error) {
	if role == "USER" {
		result, err = s.dao.CountUsers()
	} else {
		result, err = s.dao.CountInvestor()
	}
	return
}

//质押方用户注册申请批准
func (s *Service) CheckInvestorReq(c *gin.Context, id uint, stat bool) (err error) {
	entity, err := s.dao.FetchInvestorById(id)
	if stat {
		entity.AccountStatus = 1
	} else {
		_, err = s.dao.DeleteInvestorById(id)
		return
	}
	//update
	_, err = s.dao.UpdateInvestor(&entity)
	return
}

//查询所有的申请,分页要求
func (s *Service) FetchInvestorListByStatus(c *gin.Context, stat bool, pageNum, pageSize int) (investors []*schema.Investor, err error) {
	var status int
	if stat {
		status = 1
	} else {
		status = 0
	}
	list, err := s.dao.FindAllInvestorsByStatus(status, pageNum, pageSize)
	investors = make([]*schema.Investor, len(list))
	for i, item := range list {
		inv := item.Model2Schema()
		investors[i] = &inv
	}
	return
}
