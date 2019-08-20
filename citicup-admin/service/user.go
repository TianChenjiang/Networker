package service

import (
	common "citicup-admin/common"
	"citicup-admin/model"
)

type UserService struct{}

//【查询】根据userId 获取用户编号
func (service *UserService) FindOne(userId string) model.User {
	var user model.User
	orm := common.OhmEngine()
	orm.Id(userId).Get(&user)
	return user
}
