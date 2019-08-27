package service

import (
	model2 "citicup-admin/internal/model"
	"citicup-admin/schema"
	"github.com/gin-gonic/gin"
)


func (s *Service) GetAllUser(c gin.Context) (userList []*schema.User, err error) {
	list, err := s.dao.GetAllUser()
	if err != nil {
		return
	}
	bts, _ := json.Marshal(&list)
	json.Unmarshal(bts, &userList)
	return
}

func (s *Service) GetUserById(c gin.Context, id uint) (user *schema.User, err error) {
	u, err := s.dao.GetUserById(id)
	if err != nil {
		return
	}

	sche := u.Model2Schema()
	user = &sche
	return
}

func (s *Service) CreateUser(c gin.Context, userparm *schema.User) (user *schema.User, err error){
	var model = &model2.User{
		ID:       userparm.ID,
		UserName: userparm.Username,
		Password: userparm.Password,
		Email:    userparm.Email,
		Phone:    userparm.Phone,
	}
	u, err := s.dao.InsertUser(model)
	if err != nil {
		return
	}

	model_  := u.Model2Schema()
	user = &model_

	return

}

func (s *Service) UpdateUser(c gin.Context, userparm *schema.User) (err error, error error) {
	var model = &model2.User{
		ID:       userparm.ID,
		UserName: userparm.Username,
		Password: userparm.Password,
		Email:    userparm.Email,
		Phone:    userparm.Phone,
	}

	err = s.dao.UpdateUser(model)

	return
}

func (s *Service) DeleteUser(c gin.Context, id uint) (err error) {
	_, err = s.dao.DeleteUserById(id)
	return
}

func (s *Service) Check(c gin.Context, email, password string) (bool, error) {
	return s.dao.CheckAuth(email, password)
}

func (s *Service) GetUserByToken(c gin.Context, email string) (user model2.User, err error) { //todo
	user, err = s.dao.GetUserByToken(email)
	return
}









