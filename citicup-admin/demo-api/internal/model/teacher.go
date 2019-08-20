package model

import (
	xtime "github.com/CaribouW/gin-common-lib/library/time"
)

type Teacher struct {
	Id          int64          `json:"id"`
	TeacherName string         `json:"teacher_name"`
	CreateTime  xtime.Duration `json:"create_time"`
	UpdateTime  xtime.Duration `json:"update_time"`
}
