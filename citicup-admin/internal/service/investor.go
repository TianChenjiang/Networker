package service

import (
	"citicup-admin/internal/model"
	"citicup-admin/internal/web/e"
	"citicup-admin/library/util"
	"citicup-admin/schema"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func (s *Service) CheckInvestor(SwiftCode, password string) (bool, error) {
	return s.dao.CheckInvestorAuth(SwiftCode, password)
}


func (s *Service) GetInvestorByToken(c gin.Context) (investor model.Investor, code int, err error) {

	fmt.Println(c.GetHeader("Authorization")[7:])
	claim, err := util.ParseToken(c.GetHeader("Authorization")[7:])
	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		default:
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		}
		return
	}
	investor, err = s.dao.GetInvestorByToken(claim.Param)
	if err != nil {
		code = e.INTERNAL_ERROR
		return
	}
	return
}


func (s *Service) UpdateInvestorAvatar(id uint, url string) (err error) {
	err = s.dao.UpdateInvestorAvatar(id, url)
	return
}


