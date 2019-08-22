package schema

//公司 VO
type Company struct {
	ID          uint   `json:"id"`
	CompanyName string `json:"company_name"`
}

type Companies []Company
