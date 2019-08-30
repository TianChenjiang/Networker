package dao

import (
	"citicup-admin/internal/model"
	"github.com/jinzhu/gorm"
)

var user_e interface{} = &model.User{}
var companies interface{} = &model.Companies{}



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

//识别Auth
func (d *Dao) CheckAuth(email, password string) (bool, error) {
	var user model.User
	err := d.db.Select("id").Where(&model.User{Email: email, Password: password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func (d *Dao) GetUserByToken(email string) (user model.User, err error) {

	d.db.Where("email = ?", email).First(&user)
	return
}

func (d *Dao) MarkAsConcerned(companyID, userID int) (err error) {
	d.db.Model(&user_e).Related(&companies)//todo 是否指定外键

	return
}
