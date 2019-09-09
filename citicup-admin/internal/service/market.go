package service

import (
	"citicup-admin/internal/model"
	"citicup-admin/schema"
)

func (s *Service) GetAllMarket(PageNum, PageSize int) (res []model.Market, err error){
	return s.dao.GetAllMarket(PageNum, PageSize)
}

func (s *Service) InsertMarket(param schema.InsertMarketParam) (market model.Market, err error) {
	return s.dao.InsertMarket(param)
}

func (s *Service) GetMarketConditionBySymbol(symbol string) (market model.Market, err error)  {
	return s.dao.GetMarketBySymbol(symbol)
}