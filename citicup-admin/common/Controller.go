package common

import "github.com/gin-gonic/gin"

type Controller struct {
	Data interface{}
}

func (this *Controller) AjaxData(ctx *gin.Context) {
	ResultOk(ctx, this.Data)
}
func (this *Controller) Redirect(ctx *gin.Context, uri string) {
	ctx.Redirect(302, uri)
}
