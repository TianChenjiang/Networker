package model

//股票详细信息
type Stock struct {
	ID         uint   `gorm:"column:id;AUTO_INCREMENT"`
	TsCode     string `gorm:"column:ts_code"`     //TScode
	Symbol     string `gorm:"column:symbol"`      //股票代码
	Name       string `gorm:"column:name"`        //股票名称
	Area       string `gorm:"column:area"`        //所在地域
	Industry   string `gorm:"column:industry"`    //所属行业
	Fullname   string `gorm:"column:fullname"`    //股票全称
	Enname     string `gorm:"column:enname"`      //英文全称
	Market     string `gorm:"column:market"`      //市场类型
	Exchange   string `gorm:"column:exchange"`    //交易所代码
	CurrType   string `gorm:"column:curr_type"`   //交易货币
	ListStatus string `gorm:"column:list_status"` //上市状态
	ListDate   string `gorm:"column:list_date"`   //上市日期
	DelistDate string `gorm:"column:delist_date"` //退市日期
}
