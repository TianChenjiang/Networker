package bcode

import (
	"github.com/CaribouW/gin-common-lib/library/ecode"
)

// demo-api ecode interval is[1000, 2000]

var (
	OpStudentErr = ecode.New(1000, "学生操作失败")
)
