package schema

//财务信息

type Finance struct {
	Ar_to_or           string `json:"ar_to_or"`
	Ar_turn            string `json:"ar_turn"`
	Arturn_days        string `json:"arturn_days"`
	Assets_to_eqt      string `json:"assets_to_eqt"`
	Assets_turn        string `json:"assets_turn"`
	Assets_turn_days   string `json:"assets_turn_days"`
	Bps                string `json:"bps"`
	Capital_rese_ps    string `json:"capital_rese_ps"`
	Cashflow_ratio     string `json:"cashflow_ratio"`
	Current_ratio      string `json:"current_ratio"`
	Debt_to_assets     string `json:"debt_to_assets"`
	Dt_eps             string `json:"dt_eps"`
	Eps                string `json:"eps"`
	Eps_debt           string `json:"eps_debt"`
	Gross_margin       string `json:"gross_margin"`
	Grossprofit_margin string `json:"grossprofit_margin"`
	Inv_turn         string `json:"inv_turn"`
	Invturn_days     string `json:"invturn_days"`
	N_income         string `json:"n_income"`
	Netprofit_margin string `json:"netprofit_margin"`
	Netprofit_yoy    string `json:"netprofit_yoy"`
	Ocf_to_or        string `json:"ocf_to_or"`
	Ocfps            string `json:"ocfps"`
	Profit_dedt      string `json:"profit_dedt"`
	Quick_ratio      string `json:"quick_ratio"`
	Roa_dp           string `json:"roa_dp"`
	Roe_dt           string `json:"roe_dt"`
	Roe_waa          string `json:"roe_waa"`
	Salescash_to_or  string `json:"salescash_to_or"`
	Tax_rate         string `json:"tax_rate"`
	Total_revenue    string `json:"total_revenue"`
	Tr_yoy           string `json:"tr_yoy"`
	Undist_profit_ps string `json:"undist_profit_ps"`
}
