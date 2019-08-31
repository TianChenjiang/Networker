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
	bts, _ := json.Marshal(&list)
	json.Unmarshal(bts, &companies)
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
		ID:          entity.ID,
		CompanyName: entity.CompanyName,
	}
	_, err = s.dao.UpdateCompany(model_)
	return
}

func (s *Service) DeleteCompany(c *gin.Context, id uint) (err error) {
	_, err = s.dao.DeleteCompany(id)
	return
}

func (s *Service) NewCompany(c *gin.Context, entity *schema.Company) (newCompany *schema.Company, err error) {
	var model_ = &model.Company{
		ID:          entity.ID,
		CompanyName: entity.CompanyName,
	}
	com, err := s.dao.InsertCompany(model_)
	if err != nil {
		return
	}
	model_t := com.Model2Schema()
	newCompany = &model_t
	return
}
