package dao

import "citicup-admin/internal/model"

var investor_e interface{} = &model.Investor{}

//添加新的投资者
func (d *Dao) InsertInvestor(entity *model.Investor) (investor model.Investor, err error) {
	d.db.Create(entity)
	investor = *entity
	return
}