package model

import "citicup-admin/schema"

type Pledge struct {
	ID             uint    `gorm:"column:id"`               //id
	ts_code        string  `gorm:"column:ts_code"`
	CloseLine      float32 `gorm:"column:close_line"`       //平仓线
	HolderName     string  `gorm:"column:holder_name"`      //股东名称
	TotalShare     float32 `gorm:"column:total_share"`      //总股本
	PHoldingRatio  float32 `gorm:"column:p_holding_ratio"`  //占持股比（%）
	PTotalRatio    float32 `gorm:"column:p_total_ratio"`    //占总股比 (%)
	PledgeAmount   float32 `gorm:"column:pledge_amount"`    //质押数量 (TODO:单位 ? )
	Pledgor        string  `gorm:"column:pledgor"`          //质押方
	ReleaseDate    string  `gorm:"column:release_date"`     //解押日期 (TODO: 格式 ? )
	StartDate      string  `gorm:"column:start_date"`       //开始日期(TODO: 格式 ? )
	StartDateClose float32 `gorm:"column:start_date_close"` //质押日收盘价
	Status         string  `gorm:"column:status"`           //状态, "未达预警线" | "已达预警线" | "已达平仓线"TODO: 有待商讨 (是不是可以根据收盘价得出？？？)
	WarningLine    float32 `gorm:"column:warning_line"`     //预警线
}

func (p *Pledge) TableName() string {
	return "pledge"
}

func (p *Pledge) Model2Schema() (res schema.Pledge)  {
	res = schema.Pledge{
		ID:             p.ID,
		CloseLine:      p.CloseLine,
		HolderName:     p.HolderName,
		TotalShare:     p.TotalShare,
		PHoldingRatio:  p.PHoldingRatio,
		PTotalRatio:    p.PTotalRatio,
		PledgeAmount:   p.PledgeAmount,
		Pledgor:        p.Pledgor,
		ReleaseDate:    p.ReleaseDate,
		StartDate:      p.StartDate,
		StartDateClose: p.StartDateClose,
		Status:         p.Status,
		WarningLine:    p.WarningLine,
	}
	return res
}

