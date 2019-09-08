package schema

type Investor struct {
	ID       uint   `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	BankName string `json:"bank_name"`
	SwiftCode string `json:"swift_code"`
	BankEmail string `json:"bank_email"`
}

type Investors []Investor

type InvestorAuth struct {
	SwiftCode string `json:"swift_code"`
	Password string `json:"password"`
}