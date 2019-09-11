package model

type Finance struct {
	ar_to_or           string `gorm:"column:ar_to_or"`
	ar_turn            string `gorm:"column:ar_turn"`
	arturn_days        string `gorm:"column:arturn_days"`
	assets_to_eqt      string `gorm:"column:assets_to_eqt"`
	assets_turn        string `gorm:"column:assets_turn"`
	assets_turn_days   string `gorm:"column:assets_turn_days"`
	bps                string `gorm:"column:bps"`
	capital_rese_ps    string `gorm:"column:capital_rese_ps"`
	cashflow_ratio     string `gorm:"column:cashflow_ratio"`
	current_ratio      string `gorm:"column:current_ratio"`
	debt_to_assets     string `gorm:"column:debt_to_assets"`
	dt_eps             string `gorm:"column:dt_eps"`
	eps                string `gorm:"column:eps"`
	eps_debt           string `gorm:"column:eps_debt"`
	gross_margin       string `gorm:"column:gross_margin"`
	grossprofit_margin string `gorm:"column:grossprofit_margin"`
	inv_turn           string `gorm:"column:inv_turn"`
	invturn_days       string `gorm:"column:invturn_days"`
	n_income           string `gorm:"column:n_income"`
	netprofit_margin   string `gorm:"column:netprofit_margin"`
	netprofit_yoy      string `gorm:"column:netprofit_yoy"`
	ocf_to_or          string `gorm:"column:ocf_to_or"`
	ocfps              string `gorm:"column:ocfps"`
	profit_dedt        string `gorm:"column:profit_dedt"`
	quick_ratio        string `gorm:"column:quick_ratio"`
	roa_dp             string `gorm:"column:roa_dp"`
	roe_dt             string `gorm:"column:roe_dt"`
	roe_waa            string `gorm:"column:roe_waa"`
	salescash_to_or    string `gorm:"column:salescash_to_or"`
	tax_rate           string `gorm:"column:tax_rate"`
	total_revenue      string `gorm:"column:total_revenue"`
	tr_yoy             string `gorm:"column:tr_yoy"`
	undist_profit_ps   string `gorm:"column:undist_profit_ps"`
}

func (f *Finance) TableName() string {
	return "finance"
}
