package service

import (
	"citicup-admin/schema"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetAllUser(c gin.Context) (userList []*schema.User, err error) {
	_, err = s.dao.GetAllUser(c)
	return nil, nil
}
