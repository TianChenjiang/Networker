package schema

//行情数据

//市场行情		TODO:这里的市场行情是一天的？
type Market struct {
	Amount       float32 `json:"amount"`        //成交额（亿）
	Amplitude    float32 `json:"amplitude"`     //振幅 公式：(| highest - lowest | / pre_close) * 100% TODO: 这个可以根据已有数据得到的啊？
	Change       float32 `json:"change"`        //涨跌额
	PctChg       float32 `json:"pct_chg"`       //涨跌幅
	Close        float32 `json:"close"`         //收盘价
	Highest      float32 `json:"highest"`       //最高价
	ListAge      float32 `json:"list_age"`      //市龄
	Lowest       float32 `json:"lowest"`        //最低价
	Name         float32 `json:"name"`          //股票名称
	Open         float32 `json:"open"`          //开盘价
	Pb           float32 `json:"pb"`            //市净率
	Pe           float32 `json:"pe"`            //市盈率
	PreClose     float32 `json:"pre_close"`     //昨日收盘价
	Symbol       string  `json:"symbol"`        //股票代码
	TurnoverRate float32 `json:"turnover_rate"` //换手率
	Vol          float32 `json:"vol"`           //成交量
}

type InsertMarketParam struct {
	Market       Market  `json:"market"`
	CompanyID    uint    `json:"company_id"`
}