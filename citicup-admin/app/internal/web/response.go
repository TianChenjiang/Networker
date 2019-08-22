package web

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

var msgFlags = map[int]string{
	200: "ok",
	401: "method forbidden",
	403: "Unauthentic",
	404: "resource not found",
	500: "internal error",
}

func getMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}
	return msgFlags[500]
}

func (g *Gin) response(httpCode int, data interface{}) {
	if httpCode != 200 {
		g.C.JSON(httpCode, getMsg(httpCode))
	} else {
		g.C.JSON(httpCode, data)
	}
}

//200
//http.StatusOK
func (g *Gin) OK(data interface{}) {
	g.response(200, data)
}

func (g *Gin) Fail(httpCode int) {
	g.response(httpCode, getMsg(httpCode))
}
