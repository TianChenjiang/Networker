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