package dao

import "citicup-admin/internal/model"

var stock_e interface{} = &model.Stock{}

func (d *Dao) GetStockByCode(code string) (stock model.Stock, err error) {
	d.db.Model(&stock_e).Where("ts_code = ?", code).Find(&stock)
	return
}