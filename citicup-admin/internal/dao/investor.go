package dao

import (
	"citicup-admin/internal/model"
	"github.com/jinzhu/gorm"
)

var investor_e interface{} = &model.Investor{}

//添加新的投资者
func (d *Dao) InsertInvestor(entity *model.Investor) (investor model.Investor, err error) {
	d.db.Create(entity)
	investor = *entity
	return
}

func (d *Dao) GetInvestorByToken(SwiftCode string) (investor model.Investor, err error) {
	d.db.Where("swift_code = ?", SwiftCode).First(&investor)
	return
}

func (d *Dao) CheckInvestorAuth(SwiftCode, password string) (bool, error) {
	var investor model.Investor
	err := d.db.Select("id").Where(&model.Investor{SwiftCode: SwiftCode, Password: password}).First(&investor).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if investor.ID > 0 {
		return true, nil
	}

	return false, nil
}

//更新头像url
func (d *Dao) UpdateInvestorAvatar(id uint, url string) (err error) {
	d.db.Model(investor_e).Where("id = ?", id).Update("avatar", url)
	return
}

//计数质押方的用户数量
func (d *Dao) CountInvestor() (sum int, err error) {
	d.db.Model(investor_e).Count(&sum)
	return
}

//更新质押方数据
func (d *Dao) UpdateInvestor(entity *model.Investor) (investor model.Investor, err error) {
	d.db.Model(&investor_e).Where(&model.Investor{
		ID: entity.ID,
	}).Update(entity).Find(&investor)
	return
}

//根据status查询质押方
func (d *Dao) FetchInvestorByStatus(stat int) (investors []*model.Investor, err error) {
	d.db.Model(&investor_e).Where(&model.Investor{
		AccountStatus: stat,
	}).Find(&investors)
	return
}

//根据id获取用户
func (d *Dao) FetchInvestorById(id uint) (investor model.Investor, err error) {
	err = d.db.Model(&investor_e).Where(&model.Investor{ID: id}).Find(&investor).Error
	return
}

//删除investor
func (d *Dao) DeleteInvestorById(id uint) (investor model.Investor, err error) {
	err = d.db.Model(&investor_e).Where(&model.Investor{ID: id}).Delete(&investor).Error
	return
}

//获取investor列表
func (d *Dao) FindAllInvestors(pageNum, pageSize int) (investors []*model.Investor, err error) {
	err = d.db.Model(&investor_e).Offset(pageNum).Limit(pageSize).Find(&investors).Error
	return
}

//获取investor 根据状态查询
func (d *Dao) FindAllInvestorsByStatus(status int, pageNum, pageSize int) (investors []*model.Investor, err error) {
	err = d.db.Model(&investor_e).Where(&model.Investor{
		AccountStatus: status,
	}).Offset(pageNum).Limit(pageSize).Find(&investors).Error
	return
}
