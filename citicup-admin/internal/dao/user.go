package dao

import (
	"citicup-admin/internal/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

var user_e interface{} = &model.User{}
var companies []model.Company
var user model.User



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

//更新头像url
func (d *Dao) UpdateUserAvatar(id uint, url string) (err error) {
	d.db.Model(user_e).Where("id = ?", id).Update("avatar", url)
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


func (d *Dao) MarkAsConcerned(userID uint, symbol string) (err error) {
	companies = make([]model.Company, 0)
	u, _ := d.GetUserById(userID)
	d.db.Preload("Companies").First(&u, userID)
	d.db.Model(&u).Related(&companies, "Companies")
	fmt.Println(u)
	com, err  := d.GetCompanyBySymbol(symbol)
	companies = append(u.Companies, com)

	u.Companies = companies
	d.db.Model(&u).Related(&companies, "Companies").Save(&u)
	fmt.Println(u)
	return
}

func (d *Dao) UnMarkAsConcerned(userID uint, symbol string) (err error) {

	res := make([]model.Company, 0)
	u, _ := d.GetUserById(userID)
	companies, err := d.GetConcerned(0, 1000, userID) //todo

	for i := 0; i < len(companies); i++ {
		com, _  := d.GetCompanyBySymbol(symbol)
		if companies[i].ID != com.ID {
			res = append(res, companies[i])
		}
	}

	d.db.Preload("Companies").First(&u, userID)
	d.db.Model(&u).Association( "Companies").Delete(&u.Companies)
	companies = res
	u.Companies = companies
	d.db.Model(&u).Related(&companies,"Companies").Save(&u)
	fmt.Println(u)

	return
}

func (d *Dao) GetConcerned(pageNum, pageSize int, userID uint) (companyList []model.Company, err error)  {
	u, _ := d.GetUserById(userID)
	d.db.Preload("Companies").First(&u, userID)
	d.db.Model(&u).Related(&companies, "Companies").Offset(pageNum).Limit(pageSize)
	companyList = u.Companies

	return
}
