package dao

import "citicup-admin/internal/model"

var pledge_e interface{} = &model.Pledge{}
var finance_e interface{} = &model.Finance{}

func (d *Dao) GetFinanceBySymbol(symbol string) (finance model.Finance, err error)  {
	d.db.Model(finance_e).Where("ts_code = ?", symbol).First(&finance)
	return
}

func (d *Dao) GetCandleBySymbol(symbol string) (candle model.Market, err error) {
	return
}

func (d *Dao) GetPledgeBySymbol(symbol string) (pledge model.Pledge, err error) {
	//返回日期离当前最近的
	d.db.Model(pledge_e).Where("ts_code = ?", symbol).First(&pledge)
	return
}
