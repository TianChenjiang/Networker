package web

import (
	"citicup-admin/internal/publicdata"
	"citicup-admin/internal/web/e"
	"citicup-admin/library/util"
	"citicup-admin/schema"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InvestorRegister(c *gin.Context)  {
	var (
		appG = Gin{C: c}
		schema  schema.Investor
	)
	err := c.BindJSON(&schema)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.BAD_REQUEST, nil)
		return
	}

	i, err := serv.CreateInvestor(&schema)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INTERNAL_ERROR, nil)
		return
	}

	appG.OK(i)
}

func UploadRegisterInfo(c *gin.Context) {
	var (
		appG = Gin{C: c}
	)

	//获得当前银行token
	investor, code, err:= serv.GetInvestorByToken(*c)
	if err != nil {
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}


	_, image, err := c.Request.FormFile("image")
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INTERNAL_ERROR, nil)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, e.BAD_REQUEST, nil)
		return
	}

	if err := c.SaveUploadedFile(image, "library/pic/investor/"+image.Filename); err != nil { //todo 保存路径
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}

	url := "http://citicup.top/image/"+image.Filename  //todo
	err = serv.UpdateInvestorAvatar(investor.ID, url)
	appG.OK(url)
	return
}

func InvestorLogin(c *gin.Context) {
	var (
		appG = Gin{C: c}
		schema schema.InvestorAuth
	)

	err := c.BindJSON(&schema)

	isExist, err := serv.CheckInvestor(schema.SwiftCode, schema.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(schema.SwiftCode, schema.Password, publicdata.Investor)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.OK(map[string]string{
		"token": token,
	})
}
