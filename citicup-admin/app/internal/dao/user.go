package dao

import (
	"citicup-admin/internal/model"
)

var user_e interface{} = &model.User{}

//查询所有的用户
func (d *Dao) GetAllUser() (userList []*model.User, err error) {
	//查询
	d.db.Model(user_e).Find(&userList)
	return
}

//根据id获取用户
func (d *Dao) GetUserById(id uint) (user model.User, err error) {
	d.db.First(&user, id)
	return
}

//更新用户信息
func (d *Dao) UpdateUser(entity *model.User) (err error) {
	d.db.Model(user_e).Update(entity)
	return
}

//添加新的用户
func (d *Dao) InsertUser(entity *model.User) (user model.User, err error) {
	d.db.Create(entity).Find(&user)
	user = *entity
	return
}

//删除用户
func (d *Dao) DeleteUserById(id uint) (user model.User, err error) {
	d.db.Model(user_e).Where(&model.User{
		ID: id,
	}).Delete(&user)
	return
}
