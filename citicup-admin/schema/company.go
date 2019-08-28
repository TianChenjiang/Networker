package schema

//公司 VO
type Company struct {
	ID          uint   `json:"id"`
	CompanyName string `json:"companyName"`
}

type Companies []Company
