package model

import "citicup-admin/schema"

//请求
type Request struct {
	ID   uint   `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Name string `gorm:"column:name"` //名称
}

//实体转换为VO
func (r *Request) Model2Schema() (result schema.Request) {
	result = schema.Request{
		ID:   r.ID,
		Name: r.Name,
	}
	return
}
