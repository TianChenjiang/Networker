package web

import (
	"citicup-admin/internal/web/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}


type Response struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data"`
}


func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
}

//200 http.StatusOK
func (g *Gin) OK(data interface{}) {
	g.Response(http.StatusOK, e.SUCCESS, data)
}

//func (g *Gin) Fail(httpCode int) {
//	g.Response(httpCode, nil)
//}
