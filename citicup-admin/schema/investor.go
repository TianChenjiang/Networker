package schema

type Investor struct {
	ID            uint   `json:"id"`
	Password      string `json:"password"` //密码
	Email         string `json:"email"`    //邮箱
	BankName      string `json:"bank_name"`
	SwiftCode     string `json:"swift_code"`
	BankEmail     string `json:"bank_email"`
	AccountStatus int    `json:"account_status"` //账户状态,0 : 待审核, 1: 审核通过. 如果审核失败就删除账户
}

type Investors []Investor

type InvestorAuth struct {
	SwiftCode string `json:"swift_code"`
	Password  string `json:"password"`
}
