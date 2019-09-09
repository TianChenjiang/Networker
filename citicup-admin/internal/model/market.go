package model

type Market struct {
	ID           uint    `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Amount       float32 `gorm:"column:amount"`        //成交额（亿）
	Amplitude    float32 `gorm:"column:amplitude"`     //振幅 公式：(| highest - lowest | / pre_close) * 100% TODO: 这个可以根据已有数据得到的啊？
	Change       float32 `gorm:"column:change"`        //涨跌额
	PctChg       float32 `gorm:"column:pct_chg"`       //涨跌幅
	Close        float32 `gorm:"column:close"`         //收盘价
	Highest      float32 `gorm:"column:highest"`       //最高价
	ListAge      float32 `gorm:"column:list_age"`      //市龄
	Lowest       float32 `gorm:"column:lowest"`        //最低价
	Name         float32 `gorm:"column:name"`          //股票名称
	Open         float32 `gorm:"column:open"`          //开盘价
	Pb           float32 `gorm:"column:pb"`            //市净率
	Pe           float32 `gorm:"column:pe"`            //市盈率
	PreClose     float32 `gorm:"column:pre_close"`     //昨日收盘价
	TurnoverRate float32 `gorm:"column:turnover_rate"` //换手率
	Vol          float32 `gorm:"column:vol"`           //成交量
}

func (m *Market) TableName() string {
	return "market"
}


