package schema

//公司基本面信息

//公司 VO
type Company struct {
	ID            uint   `json:"id"`
	TsCode        string `json:"ts_code"`        //股票代码
	Chairman      string `json:"chairman"`       //法人代表
	Manager       string `json:"manager"`        //总经理
	Secretary     string `json:"secretary"`      //董秘
	RegCapital    string `json:"reg_capital"`    //注册资本
	SetupDate     string `json:"setup_date"`     //创立日期
	Province      string `json:"province"`       //所在省份
	City          string `json:"city"`           //所在城市
	Introduction  string `json:"introduction"`   //公司介绍
	Website       string `json:"website"`        //公司主页URL
	Email         string `json:"email"`          //电子邮件
	Office        string `json:"office"`         //办公室
	Employees     uint   `json:"employees"`      //员工人数
	MainBusiness  string `json:"main_business"`  //主要业务及产品
	BusinessScope string `json:"business_scope"` //经营范围
	Market        Market `json:"market"`
	MarketID      uint   `json:"market_id"` //市场行情与公司一对一
}

type Companies []Company

//预测结果
type PredictionResult struct {
	BurstRate float64 `json:"burst_rate"` //爆仓概率
	Name      string  `json:"name"`       //股票名称 (TODO: 可以进行关联)
	RiskLevel float64 `json:"risk_level"` //risk_level: "高" | "中" | "低" TODO: 需要给出具体量化指标
	Symbol    string  `json:"symbol"`     //股票代码 (TODO: 可以进行关联)
}

//股票详细信息
type Stock struct {
	Symbol     string `json:"symbol"`      //股票代码
	Name       string `json:"name"`        //股票名称
	Area       string `json:"area"`        //所在地域
	Industry   string `json:"industry"`    //所属行业
	Fullname   string `json:"fullname"`    //股票全称
	Enname     string `json:"enname"`      //英文全称
	Market     string `json:"market"`      //市场类型
	Exchange   string `json:"exchange"`    //交易所代码
	CurrType   string `json:"curr_type"`   //交易货币
	ListStatus string `json:"list_status"` //上市状态
	ListDate   string `json:"list_date"`   //上市日期
	DelistDate string `json:"delist_date"` //退市日期
}
