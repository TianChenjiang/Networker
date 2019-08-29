package schema

//质押详情
type Pledge struct {
	ID             uint    `json:"id"`               //id
	CloseLine      float32 `json:"close_line"`       //平仓线
	HolderName     string  `json:"holder_name"`      //股东名称
	PHoldingRatio  float32 `json:"p_holding_ratio"`  //占持股比（%）
	PTotalRatio    float32 `json:"p_total_ratio"`    //占总股比 (%)
	PledgeAmount   float32 `json:"pledge_amount"`    //质押数量 (TODO:单位 ? )
	Pledgor        string  `json:"pledgor"`          //质押方
	ReleaseDate    string  `json:"release_date"`     //解押日期 (TODO: 格式 ? )
	StartDate      string  `json:"start_date"`       //开始日期(TODO: 格式 ? )
	StartDateClose float32 `json:"start_date_close"` //质押日收盘价
	Status         string  `json:"status"`           //状态, "未达预警线" | "已达预警线" | "已达平仓线"TODO: 有待商讨 (是不是可以根据收盘价得出？？？)
	WarningLine    float32 `json:"warning_line"`     //预警线
}


