package common

import "math"

//分页对象
type Pager struct {
	Page      int64   `form:"page"  json:"page"`           //当前页
	PageSize  int64   `form:"pageSize"  json:"pageSize"`   //每页条数
	Total     int64   `form:"total"  json:"total"`         //总条数
	PageCount int64   `form:"pageCount"  json:"pageCount"` //总页数
	Nums      []int64 `form:"nums"  json:"nums"`           //分页序数
	NumsCount int64   `form:"numsCount"  json:"numsCount"` //总页序数
}

func CreatePager(page, pagesize, total int64) *Pager {
	if page < 1 {
		page = 1
	}
	if pagesize < 1 {
		pagesize = 10
	}

	page_count := math.Ceil(float64(total) / float64(pagesize))

	pager := new(Pager)
	pager.Page = page
	pager.PageSize = pagesize
	pager.Total = total
	pager.PageCount = int64(page_count)
	pager.NumsCount = 7
	pager.setNums()
	return pager
}

func (this *Pager) setNums() {
	this.Nums = []int64{}
	if this.PageCount == 0 {
		return
	}

	half := math.Floor(float64(this.NumsCount) / float64(2))
	begin := this.Page - int64(half)
	if begin < 1 {
		begin = 1
	}

	end := begin + this.NumsCount - 1
	if end >= this.PageCount {
		begin = this.PageCount - this.NumsCount + 1
		if begin < 1 {
			begin = 1
		}
		end = this.PageCount
	}

	for i := begin; i <= end; i++ {
		this.Nums = append(this.Nums, i)
	}
}
