package schema

//行情数据

//市场行情		TODO:这里的市场行情是一天的？
type Market struct {
	ID       uint    `json:"id"`
	Tscode   string  `json:"column:ts_code"` //股票代码
	TradeDate string  `json:"trade_date"` //交易日期
	Open     float32 `json:"open"`           //开盘价
	Amount   float32 `json:"amount"`         //成交额（亿）
	Change   float32 `json:"change"`         //涨跌额
	PctChg   float32 `json:"pct_chg"`        //涨跌幅
	Close    float32 `json:"close"`          //收盘价
	Highest  float32 `json:"high"`           //最高价
	Lowest   float32 `json:"low"`            //最低价
	PreClose float32 `json:"pre_close"`      //昨日收盘价
	Vol      float32 `json:"vol"`            //成交量
}

type InsertMarketParam struct {
	Market    Market `json:"market"`
	CompanyID uint   `json:"company_id"`
}
