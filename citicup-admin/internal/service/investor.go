package service

import (
	"citicup-admin/internal/model"
	"citicup-admin/schema"
)

func (s *Service) CreateInvestor(investorParm *schema.Investor) (investor *schema.Investor, err error) {
	var model = &model.Investor{
		Password:  investorParm.Password,
		Email:     investorParm.Email,
		BankName:  investorParm.BankName,
		SwiftCode: investorParm.SwiftCode,
		BankEmail: investorParm.BankEmail,
	}

	i, err := s.dao.InsertInvestor(model)
	if err != nil {
		return
	}

	model_  := i.Model2Schema()
	investor = &model_

	return
	
}
