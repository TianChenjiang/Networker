package service

import (
	"citicup-admin/internal/model"
	"citicup-admin/internal/web/e"
	"citicup-admin/library/util"
	"citicup-admin/schema"
	"fmt"
	"github.com/dgrijalva/jwt-go"
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
	var model = &model.User{
		ID:       userparm.ID,
		Username: userparm.Username,
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
	var model = &model.User{
		ID:       userparm.ID,
		Username: userparm.Username,
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

//func (s *Service) GetUserByToken(c gin.Context, email string) (user model2.User, err error) { //todo
//	user, err = s.dao.GetUserByToken(email)
//	return
//}

func (s *Service) GetUserByToken(c gin.Context) (user model.User, code int, err error) {

	fmt.Println(c.GetHeader("Authorization")[7:])
	claim, err := util.ParseToken(c.GetHeader("Authorization")[7:])
	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		default:
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		}
		return
	}
	user, err = s.dao.GetUserByToken(claim.Email)
	if err != nil {
		code = e.INTERNAL_ERROR
		return
	}
	return
}

func (s *Service) ChangePassword(c gin.Context, password string, userparm model.User) (err error) {
	var model = &model.User{
		ID:       userparm.ID,
		Username: userparm.Username,
		Password: password,
		Email:    userparm.Email,
		Phone:    userparm.Phone,
	}
	err = s.dao.UpdateUser(model)
	return
}

func (s *Service) MarkAsConcerned(c gin.Context, ID uint, symbol string) (err error) {
	err = s.dao.UpdateUser()


}


