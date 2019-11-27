package service

import "citicup-admin/internal/model"

func (s *Service) GetFinanceBySymbol(symbol string) (finance model.Finance, err error)  {
	return s.dao.GetFinanceBySymbol(symbol)
}

func (s *Service) GetCandleBySymbol(symbol string) (candle model.Market, err error) {
	return
}

func (s *Service) GetPledgeBySymbol(symbol string) (pledge model.Pledge, err error) {
	return s.dao.GetPledgeBySymbol(symbol)
}
