package app

import (
	"github.com/gin-gonic/gin"
	"networker/backend/src/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code   int           `json:"code"`
	Msg    string        `json:"msg"`
	Data   interface{}   `json:"data"`
}

func (g *Gin) Response(httpCode int, errorCode int, data interface{})  {
	g.C.JSON(httpCode, Response{
		Code: httpCode,
		Msg:  e.GetMsg(errorCode),
		Data: data,
	})
}
