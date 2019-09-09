package schema

//分页参数
type Pagination struct {
	PageNum int `json:"page_num"`
	PageSize int `json:"page_size"`
}

//默认分页参数
func DefaultPagination() Pagination {
	return Pagination{
		PageNum:  0,
		PageSize: 5,
	}
}