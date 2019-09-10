package service

import (
	"citicup-admin/internal/model"
	"citicup-admin/schema"
	"github.com/gin-gonic/gin"
)

//获取所有的公司
func (s *Service) GetAllCompanies(c *gin.Context) (companies []*schema.Company, err error) {
	list, err := s.dao.GetAllCompanies()
	if err != nil {
		return
	}
	companies = make([]*schema.Company, len(list))
	for i, item := range list {
		c := item.Model2Schema()
		companies[i] = &c
	}
	return
}

//分页获取所有公司
func (s *Service) GetAllCompaniesPaging(c *gin.Context, pageNum, pageSize int) (companies []*schema.Company, err error) {
	list, err := s.dao.GetAllCompaniesPaging(pageNum, pageSize)
	if err != nil {
		return
	}
	companies = make([]*schema.Company, len(list))
	for i, item := range list {
		c := item.Model2Schema()
		companies[i] = &c
	}
	return
}

//根据Id获取公司
func (s *Service) GetCompanyById(c *gin.Context, id uint) (company *schema.Company, err error) {
	entity, err := s.dao.GetCompanyById(id)
	if err != nil {
		return
	}
	var schema_ = entity.Model2Schema()
	company = &schema_
	return
}

//更新公司信息
func (s *Service) UpdateCompany(c *gin.Context, entity *schema.Company) (err error) {
	var model_ = &model.Company{
		ID:            entity.ID,
		CompanyName:   entity.CompanyName,
		Chairman:      entity.Chairman,
		Manager:       entity.Manager,
		Secretary:     entity.Secretary,
		RegCapital:    entity.RegCapital,
		SetupDate:     entity.SetupDate,
		Province:      entity.Province,
		City:          entity.City,
		Introduction:  entity.Introduction,
		Website:       entity.Website,
		Email:         entity.Email,
		Office:        entity.Office,
		Employees:     entity.Employees,
		MainBusiness:  entity.MainBusiness,
		BusinessScope: entity.BusinessScope,
	}
	_, err = s.dao.UpdateCompany(model_)
	return
}

//删除公司信息
func (s *Service) DeleteCompany(c *gin.Context, id uint) (err error) {
	_, err = s.dao.DeleteCompany(id)
	return
}

func (s *Service) NewCompany(c *gin.Context, entity *schema.Company) (newCompany *schema.Company, err error) {
	var model_ = &model.Company{
		ID:            entity.ID,
		CompanyName:   entity.CompanyName,
		Chairman:      entity.Chairman,
		Manager:       entity.Manager,
		Secretary:     entity.Secretary,
		RegCapital:    entity.RegCapital,
		SetupDate:     entity.SetupDate,
		Province:      entity.Province,
		City:          entity.City,
		Introduction:  entity.Introduction,
		Website:       entity.Website,
		Email:         entity.Email,
		Office:        entity.Office,
		Employees:     entity.Employees,
		MainBusiness:  entity.MainBusiness,
		BusinessScope: entity.BusinessScope,
	}
	com, err := s.dao.InsertCompany(model_)
	if err != nil {
		return
	}
	model_t := com.Model2Schema()
	newCompany = &model_t
	return
}

func (s *Service) GetMarket(companyID uint) (market model.Market, err error) {
	market, err = s.dao.GetMarket(companyID)
	return
}
