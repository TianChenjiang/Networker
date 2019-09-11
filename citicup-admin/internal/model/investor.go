package model

import "citicup-admin/schema"

//用户实体
type Investor struct {
	ID            uint   `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Password      string `gorm:"column:password"`
	Email         string `gorm:"column:email"`
	BankName      string `gorm:"column:bank_name"`
	SwiftCode     string `gorm:"column:swift_code"`
	BankEmail     string `gorm:"column:bank_email"`
	Avatar        string `gorm:"column:avatar"`
	AccountStatus int    `gorm:"column:account_status"` //账户状态,0 : 待审核, 1: 审核通过. 如果审核失败就删除账户
}

func (i *Investor) TableName() string {
	return "investor"
}

func (i *Investor) Model2Schema() (result schema.Investor) {
	result = schema.Investor{
		ID:            i.ID,
		Password:      i.Password,
		Email:         i.Email,
		BankName:      i.BankName,
		SwiftCode:     i.SwiftCode,
		BankEmail:     i.BankEmail,
		AccountStatus: i.AccountStatus,
	}
	return
}

type Investors []Investor

func (i Investors) Model2Schema() (result []schema.Investor) {
	result = make([]schema.Investor, len(i))
	for j, item := range i {
		result[j] = item.Model2Schema()
	}
	return
}
