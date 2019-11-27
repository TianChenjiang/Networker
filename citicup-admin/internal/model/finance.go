package model

import "citicup-admin/schema"

type Finance struct {
	ts_code            string `gorm:column:ts_code`
	Ar_to_or           string `gorm:"column:ar_to_or"`
	Ar_turn            string `gorm:"column:ar_turn"`
	Arturn_days        string `gorm:"column:arturn_days"`
	Assets_to_eqt      string `gorm:"column:assets_to_eqt"`
	Assets_turn        string `gorm:"column:assets_turn"`
	Assets_turn_days   string `gorm:"column:assets_turn_days"`
	Bps                string `gorm:"column:bps"`
	Capital_rese_ps    string `gorm:"column:capital_rese_ps"`
	Cashflow_ratio     string `gorm:"column:cashflow_ratio"`
	Current_ratio      string `gorm:"column:current_ratio"`
	Debt_to_assets     string `gorm:"column:debt_to_assets"`
	Dt_eps             string `gorm:"column:dt_eps"`
	Eps                string `gorm:"column:eps"`
	Eps_debt           string `gorm:"column:eps_debt"`
	Gross_margin       string `gorm:"column:gross_margin"`
	Grossprofit_margin string `gorm:"column:grossprofit_margin"`
	Inv_turn         string `gorm:"column:inv_turn"`
	Invturn_days     string `gorm:"column:invturn_days"`
	N_income         string `gorm:"column:n_income"`
	Netprofit_margin string `gorm:"column:netprofit_margin"`
	Netprofit_yoy    string `gorm:"column:netprofit_yoy"`
	Ocf_to_or        string `gorm:"column:ocf_to_or"`
	Ocfps            string `gorm:"column:ocfps"`
	Profit_dedt      string `gorm:"column:profit_dedt"`
	Quick_ratio      string `gorm:"column:quick_ratio"`
	Roa_dp           string `gorm:"column:roa_dp"`
	Roe_dt           string `gorm:"column:roe_dt"`
	Roe_waa          string `gorm:"column:roe_waa"`
	Salescash_to_or  string `gorm:"column:salescash_to_or"`
	Tax_rate         string `gorm:"column:tax_rate"`
	Total_revenue    string `gorm:"column:total_revenue"`
	Tr_yoy           string `gorm:"column:tr_yoy"`
	Undist_profit_ps string `gorm:"column:undist_profit_ps"`
}

func (f *Finance) TableName() string {
	return "finance"
}

func (f *Finance) Model2Schema() (res schema.Finance){
	res = schema.Finance{
		Ar_to_or:           f.Ar_to_or,
		Ar_turn:            f.Ar_turn,
		Arturn_days:        f.Arturn_days,
		Assets_to_eqt:      f.Assets_to_eqt,
		Assets_turn:        f.Assets_turn,
		Assets_turn_days:   f.Assets_turn_days,
		Bps:                f.Bps,
		Capital_rese_ps:    f.Capital_rese_ps,
		Cashflow_ratio:     f.Cashflow_ratio,
		Current_ratio:      f.Current_ratio,
		Debt_to_assets:     f.Debt_to_assets,
		Dt_eps:             f.Dt_eps,
		Eps:                f.Eps,
		Eps_debt:           f.Eps_debt,
		Gross_margin:       f.Gross_margin,
		Grossprofit_margin: f.Grossprofit_margin,
		Inv_turn:           f.Inv_turn,
		Invturn_days:       f.Invturn_days,
		N_income:           f.N_income,
		Netprofit_margin:   f.Netprofit_margin,
		Netprofit_yoy:      f.Netprofit_yoy,
		Ocf_to_or:          f.Ocf_to_or,
		Ocfps:              f.Ocfps,
		Profit_dedt:        f.Profit_dedt,
		Quick_ratio:        f.Quick_ratio,
		Roa_dp:             f.Roa_dp,
		Roe_dt:             f.Roe_dt,
		Roe_waa:            f.Roe_waa,
		Salescash_to_or:    f.Salescash_to_or,
		Tax_rate:           f.Tax_rate,
		Total_revenue:      f.Total_revenue,
		Tr_yoy:             f.Tr_yoy,
		Undist_profit_ps:   f.Undist_profit_ps,
	}
	return
}
