package dao

import "citicup-admin/internal/model"


var market_e interface{} = &model.Market{}

func (d *Dao) GetAllMarket(PageNum, PageSize int) (res []model.Market, err error) {
	d.db.Model(market_e).Find(&res)
	return
}
