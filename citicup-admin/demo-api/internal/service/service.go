package service

import (
	"citicup-admin/demo-api/internal/conf"
	"citicup-admin/demo-api/internal/dao"
	"context"
)

// Service struct
type Service struct {
	c   *conf.Config
	dao *dao.Dao
}

// New init
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
	}
	return
}

// Ping ping the resource.
func (s *Service) Ping(c context.Context) (err error) {
	// TODO
	return
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}
