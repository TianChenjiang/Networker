package service

import (
	"citicup-admin/internal/model"
	"citicup-admin/schema"
)

func (s *Service) GetAllMarket(PageNum, PageSize int) (res []string, totalNum int, err error) {
	markets, err := s.dao.GetAllMarket(PageNum, PageSize)
	for i := 0; i < len(markets); i++ {
		res = append(res, markets[i].Tscode)
	}
	totalNum = len(res)
	return
}


func (s *Service) InsertMarket(param schema.InsertMarketParam) (market model.Market, err error) {
	return s.dao.InsertMarket(param)
}

func (s *Service) GetMarketConditionBySymbol(userID uint, symbol string) (market model.Market, isConcerned bool, name string, err error) {
	return s.dao.GetMarketBySymbol(userID, symbol)
}
