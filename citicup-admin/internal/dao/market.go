package dao

import (
	"citicup-admin/internal/model"
	"citicup-admin/schema"
)


var market_e interface{} = &model.Market{}

func (d *Dao) GetAllMarket(PageNum, PageSize int) (res []model.Market, err error) {
	d.db.Model(market_e).Offset(PageNum).Limit(PageSize).Find(&res)
	return
}

func (d *Dao) InsertMarket(param schema.InsertMarketParam) (market model.Market, err error)  {
	d.db.Model(market_e).Create(param.Market).Find(&market)
	company, _ := d.GetCompanyById(param.CompanyID)
	company.Market = market
	company.MarketID = market.ID

	d.db.Preload("market_id")
	d.db.Save(&company)

	return
}

func (d *Dao) GetMarketBySymbol(userID uint, symbol string) (market model.Market, isConcerned bool, name string, err error)  {
	d.db.Where("ts_code = ?", symbol).First(&market)
	companyList, err := d.GetConcerned(0, 1000, userID)
	company, err := d.GetCompanyBySymbol(symbol)
	stock, err := d.GetStockByCode(symbol)
	name = stock.Name


	isConcerned = false
	for i := 0; i < len(companyList); i++  {
		if company == companyList[i] {
			isConcerned = true
			break
		}
	}
	return
}