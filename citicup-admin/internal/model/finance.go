package model

import "citicup-admin/schema"

type Finance struct {
	ts_code          string `gorm:column:ts_code`
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

func (f *Finance) Model2Schema() (res schema.Finance){
	res = schema.Finance{
		Ar_to_or:           f.ar_to_or,
		Ar_turn:            f.ar_turn,
		Arturn_days:        f.arturn_days,
		Assets_to_eqt:      f.assets_to_eqt,
		Assets_turn:        f.assets_turn,
		Assets_turn_days:   f.assets_turn_days,
		Bps:                f.bps,
		Capital_rese_ps:    f.capital_rese_ps,
		Cashflow_ratio:     f.cashflow_ratio,
		Current_ratio:      f.current_ratio,
		Debt_to_assets:     f.debt_to_assets,
		Dt_eps:             f.dt_eps,
		Eps:                f.eps,
		Eps_debt:           f.eps_debt,
		Gross_margin:       f.gross_margin,
		Grossprofit_margin: f.grossprofit_margin,
		Inv_turn:           f.inv_turn,
		Invturn_days:       f.invturn_days,
		N_income:           f.n_income,
		Netprofit_margin:   f.netprofit_margin,
		Netprofit_yoy:      f.netprofit_yoy,
		Ocf_to_or:          f.ocf_to_or,
		Ocfps:              f.ocfps,
		Profit_dedt:        f.profit_dedt,
		Quick_ratio:        f.quick_ratio,
		Roa_dp:             f.roa_dp,
		Roe_dt:             f.roe_dt,
		Roe_waa:            f.roe_waa,
		Salescash_to_or:    f.salescash_to_or,
		Tax_rate:           f.tax_rate,
		Total_revenue:      f.total_revenue,
		Tr_yoy:             f.tr_yoy,
		Undist_profit_ps:   f.undist_profit_ps,
	}
	return
}
