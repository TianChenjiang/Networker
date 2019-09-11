package service

import "citicup-admin/internal/model"

func (s *Service) GetFinanceBySymbol(symbol string) (finace model.Finance, err error)  {
	return s.dao.GetFinanceBySymbol(symbol)
}