package dao

import (
	"citicup-admin/internal/model"
	"github.com/gin-gonic/gin"
)

var model_e interface{} = &model.User{}

//查询所有的用户
func (d *Dao) GetAllUser(c gin.Context) (userList []*model.User, err error) {
	//查询
	d.db.Model(model_e).Find(&userList)
	return
}

//根据id获取用户
func (d *Dao) GetUserById(c gin.Context, id uint) (user *model.User, err error) {
	d.db.Model(model_e).Where(&model.User{ID: id}).Find(user)
	return
}

//更新用户信息
func (d *Dao) UpdateUser(c gin.Context, entity *model.User) (err error) {
	d.db.Model(model_e).Update(entity)
	return
}

//添加新的用户
func (d *Dao) InsertUser(c gin.Context, entity *model.User) (user *model.User, err error) {
	d.db.Model(model_e).Save(entity)
	user = entity
	return
}
