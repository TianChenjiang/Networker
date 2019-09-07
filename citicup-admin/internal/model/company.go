package model

import (
	"citicup-admin/schema"
)

//model 公司实体
type Company struct {
	ID          uint   `gorm:"column:id;AUTO_INCREMENT"`
	CompanyName string `gorm:"column:company_name"`
	Symbol      string `gorm:"column:symbol""`
	//Users       []*User `gorm:"many2many:user_company;"`
}

type Companies []Company

func (c *Company) TableName() string {
	return "company"
}

//单个资源转换
func (c *Company) Model2Schema() (result schema.Company) {
	result = schema.Company{
		ID:          c.ID,
		CompanyName: c.CompanyName,
	}
	return
}

func (c Companies) Model2Schema(result schema.Companies) {
	result = make(schema.Companies, len(c))
	for i, item := range c {
		result[i] = item.Model2Schema()
	}
	return
}

