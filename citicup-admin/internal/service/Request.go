package service

import "citicup-admin/schema"

func (s *Service) GetAllRequests(PageNum, PageSize int) (result []*schema.Request, err error) {
	list, err := s.dao.GetAllRequests()
	if err != nil {
		return
	}
	result = make([]*schema.Request, len(list))
	for i, item := range list {
		c := item.Model2Schema()
		result[i] = &c
	}
	return
}
