package model

import (
	"citicup-admin/schema"
)

//model 公司实体
type Company struct {
	ID            uint   `gorm:"column:id;AUTO_INCREMENT"`
	CompanyName   string `gorm:"column:company_name"`
	Symbol        string `gorm:"column:symbol"` //股票代码
	Market        Market
	MarketID      uint                                  //市场行情与公司一对一
	Chairman      string `gorm:"column:chairman"`       //法人代表
	Manager       string `gorm:"column:manager"`        //总经理
	Secretary     string `gorm:"column:secretary"`      //董秘
	RegCapital    string `gorm:"column:reg_capital"`    //注册资本
	SetupDate     string `gorm:"column:setup_date"`     //创立日期
	Province      string `gorm:"column:province"`       //所在省份
	City          string `gorm:"column:city"`           //所在城市
	Introduction  string `gorm:"column:introduction"`   //公司介绍
	Website       string `gorm:"column:website"`        //公司主页URL
	Email         string `gorm:"column:email"`          //电子邮件
	Office        string `gorm:"column:office"`         //办公室
	Employees     uint   `gorm:"column:employees"`      //员工人数
	MainBusiness  string `gorm:"column:main_business"`  //主要业务及产品
	BusinessScope string `gorm:"column:business_scope"` //经营范围
}

type Companies []Company

func (c *Company) TableName() string {
	return "company"
}

//单个资源转换
func (c *Company) Model2Schema() (result schema.Company) {
	result = schema.Company{
		ID:            c.ID,
		CompanyName:   c.CompanyName,
		Chairman:      c.Chairman,
		Manager:       c.Manager,
		Secretary:     c.Secretary,
		RegCapital:    c.RegCapital,
		SetupDate:     c.SetupDate,
		Province:      c.Province,
		City:          c.City,
		Introduction:  c.Introduction,
		Website:       c.Website,
		Email:         c.Email,
		Office:        c.Office,
		Employees:     c.Employees,
		MainBusiness:  c.MainBusiness,
		BusinessScope: c.BusinessScope,
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
