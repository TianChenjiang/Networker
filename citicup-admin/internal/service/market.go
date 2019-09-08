package service

import "citicup-admin/internal/model"

func (s *Service) GetAllMarket(PageNum, PageSize int) (res []model.Market, err error){
	return s.dao.GetAllMarket(PageNum, PageSize)
}
