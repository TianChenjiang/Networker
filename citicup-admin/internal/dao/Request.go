package dao

import "citicup-admin/internal/model"

var request_e interface{} = &model.Request{}

//获取所有的申请
func (d *Dao) GetAllRequests() (requests []*model.Request, err error) {
	d.db.Model(request_e).Find(&requests)
	return
}
