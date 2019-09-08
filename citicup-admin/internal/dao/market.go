package dao

import (
	"citicup-admin/internal/model"
	"citicup-admin/schema"
)


var market_e interface{} = &model.Market{}

func (d *Dao) GetAllMarket(PageNum, PageSize int) (res []model.Market, err error) {
	d.db.Model(market_e).Find(&res)
	return
}

func (d *Dao) InsertMarket(param schema.InsertMarketParam) (market model.Market, err error)  {
	d.db.Model(market_e).Create(&param.Market).Find(&market)
	company, _ := d.GetCompanyById(param.CompanyID)
	company.Market = market

	d.db.Preload("market_id")
	d.db.Model(&market).Related(&company, "market_id").Save(&company) //todo company列未更新

	return
}