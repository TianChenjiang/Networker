package schema

//财务数据相关

//利润表
type Income struct {
	AnnDate      string  `json:"ann_date"`      //公告日期
	FAnnDate     string  `json:"f_ann_date"`    //实际公告日期
	EndDate      string  `json:"end_date"`      //报告期
	ReportType   string  `json:"report_type"`   //报告类型
	CompType     string  `json:"comp_type"`     //公司类型
	BasicEps     float32 `json:"basic_eps"`     //基本每股收益
	DilutedEps   float32 `json:"diluted_eps"`   //稀释每股收益
	TotalRevenue float32 `json:"total_revenue"` //营业总收入
	Revenue      float32 `json:"revenue"`       //营业收入
	IntIncome    float32 `json:"int_income"`    //利息收入
	PremEarned   float32 `json:"prem_earned"`   //已赚保费
}

//现金流量表
type CashFlow struct {
}

//财务指标数据
type FinaIndicator struct {
}
