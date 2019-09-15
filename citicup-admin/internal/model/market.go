package model

import "citicup-admin/schema"

//日线行情
type Market struct {
	ID        uint    `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Tscode    string  `gorm:"column:ts_code"`
	Open      float32 `gorm:"column:open"`       //开盘价
	TradeDate string  `gorm:"column:trade_date"` //交易日期
	Amount    float32 `gorm:"column:amount"`     //成交额（亿）
	Change    float32 `gorm:"column:change"`     //涨跌额
	PctChg    float32 `gorm:"column:pct_chg"`    //涨跌幅
	Close     float32 `gorm:"column:close"`      //收盘价
	Highest   float32 `gorm:"column:high"`       //最高价
	Lowest    float32 `gorm:"column:low"`        //最低价
	PreClose  float32 `gorm:"column:pre_close"`  //昨日收盘价
	Vol       float32 `gorm:"column:vol"`        //成交量
}
type Markets []Market

func (m *Market) Model2Schema() (result schema.Market) {
	result = schema.Market{
		ID:        m.ID,
		Tscode:    m.Tscode,
		Amount:    m.Amount,
		TradeDate: m.TradeDate,
		Change:    m.Change,
		PctChg:    m.PctChg,
		Close:     m.Close,
		Highest:   m.Highest,
		Lowest:    m.Lowest,
		Open:      m.Open,
		PreClose:  m.PreClose,
		Vol:       m.Vol,
	}
	return
}

func (m *Market) TableName() string {
	return "market"
}
