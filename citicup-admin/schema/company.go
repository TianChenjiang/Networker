package schema

//公司 VO
type Company struct {
	ID          uint   `json:"id"`
	CompanyName string `json:"companyName"`
}

type Companies []Company

//预测结果
type PredictionResult struct {
	BurstRate float64 `json:"burst_rate"` //爆仓概率
	Name      string  `json:"name"`       //股票名称 (TODO: 可以进行关联)
	RiskLevel float64 `json:"risk_level"` //risk_level: "高" | "中" | "低" TODO: 需要给出具体量化指标
	Symbol    string  `json:"symbol"`     //股票代码 (TODO: 可以进行关联)
}

