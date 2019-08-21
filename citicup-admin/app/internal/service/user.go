package service

import (
	"citicup-admin/schema"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetAllUser(c gin.Context) (userList []*schema.User, err error) {
	list, err := s.dao.GetAllUser(c)
	if err != nil {
		return
	}
	userList = make([]*schema.User, 0)
	for _, item := range list {
		u := item.Model2Schema()
		userList = append(userList, &u)
	}
	return
}

func (s *Service) GetUserById(c gin.Context, id uint) (user *schema.User, err error) {
	u, err := s.dao.GetUserById(c, id)
	if err != nil {
		return
	}
	//Get the users
	sche := u.Model2Schema()
	user = &sche
	return
}

func (s *Service) UpdateUser(c gin.Context, schemaEntity *schema.User) (err error) {
	return
}
