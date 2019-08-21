package service

import (
	"citicup-admin/internal/config"
	"citicup-admin/internal/dao"
	"context"
)

// Service struct
type Service struct {
	c   *config.Config
	dao *dao.Dao
}

// New init
func New(c *config.Config) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
	}
	return
}

// Ping ping the resource.
func (s *Service) Ping(c context.Context) (err error) {
	return
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}
