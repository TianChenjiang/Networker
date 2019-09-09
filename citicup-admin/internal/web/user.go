package web

import (
	"citicup-admin/internal/publicdata"
	"citicup-admin/internal/web/e"
	"citicup-admin/library/util"
	"citicup-admin/schema"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)


func GetUserList(c *gin.Context) {
	var (
		appG = Gin{C: c}
	)
	fmt.Println(c.GetHeader("User-Agent"))
	list, err := serv.GetAllUser(*c)
	if err != nil {
		log.Print("err")
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USERLIST, nil)
		return
	}
	appG.OK(list)
}


func GetUserByID(c *gin.Context) {
	var (
		appG  = Gin{C: c}
		id, _ = strconv.Atoi(c.Param("id"))
	)

	user, err := serv.GetUserById(*c, uint(id))
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USER, nil)
		return
	}
	appG.OK(user)
}

func Register(c *gin.Context) {
	var (
		appG = Gin{C: c}
		schema  schema.User
	)
	err := c.BindJSON(&schema)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_CREATE_USER, nil)
		return
	}

	u, err := serv.CreateUser(*c, &schema)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_USER, nil)
		return
	}

	appG.OK(u)
}

func EditUser(c *gin.Context) {
	var (
		appG = Gin{C: c}
		schema  schema.User
	)
	err := c.Bind(&schema)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_EDIT_USER, nil)
		return
	}

	u, err := serv.UpdateUser(*c, &schema)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_USER, nil)
		return
	}

	appG.OK(u)
}

func DeleteUser(c *gin.Context) {
	var (
		appG = Gin{C: c}
		id, _ = strconv.Atoi(c.Param("id"))
	)

	err := serv.DeleteUser(*c, uint(id))
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_USER, nil)
		return
	}
	appG.OK(err) //todo
}

// @Summary user login
// @Produce  json
// @Param email query string true "email"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/users/token/:token [get]
func UserLogin(c *gin.Context) {
	var (
		appG = Gin{C: c}
		schema schema.UserAuth
	)

	err := c.BindJSON(&schema)


	isExist, err := serv.CheckUser(*c, schema.Email, schema.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(schema.Email, schema.Password, publicdata.Investor)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.OK(map[string]string{
		"token": token,
	})
}

func ChangePassword(c *gin.Context)  {
	var (
		appG = Gin{C: c}
		schema schema.ChangePasswordSchema
	)

	c.BindJSON(&schema)

	user, code, err:= serv.GetUserByToken(*c)
	if err != nil {
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}

	if schema.Password == user.Password {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHANGE_PASSWORD_SAME, nil)
		return
	}
	err = serv.ChangePassword(*c, schema.Password, user)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHANGE_PASSWORD, nil)
		return
	}
	appG.OK("成功修改密码")
}

func UploadAvatar(c *gin.Context) {
	var (
		appG = Gin{C: c}
	)

	//获得当前用户token
	user, code, err:= serv.GetUserByToken(*c)
	if err != nil {
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}

	//claim, err := util.ParseToken(c.GetHeader("Authorization")[7:])
	//role := claim.Role
	//if role != publicdata.User {
	//	return
	//}


	_, image, err := c.Request.FormFile("image")
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INTERNAL_ERROR, nil)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, e.BAD_REQUEST, nil)
		return
	}

	if err := c.SaveUploadedFile(image, "library/pic/user/"+image.Filename); err != nil { //todo 保存路径
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}
	//更新用户的avatar属性
	url := "http://citicup.top/image/"+image.Filename  //todo
	err = serv.UpdateUserAvatar(user.ID, url)
	appG.OK(url)
	return
}

func MarkAsConcerned(c *gin.Context) {
	var (
		appG = Gin{C: c}
		symbol = c.Query("symbol")
	)
	user, code, err:= serv.GetUserByToken(*c)
	if err != nil {
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}

	err = serv.MarkAsConcerned(*c, user.ID, symbol)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INTERNAL_ERROR, nil)
	}

	appG.OK("标记成功")
}


func CancelMarkAsConcerned(c *gin.Context) {
	var (
		appG = Gin{C: c}
		symbol = c.Query("symbol")
	)
	user, code, err:= serv.GetUserByToken(*c)
	if err != nil {
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}

	err = serv.UnMarkAsConcerned(*c, user.ID, symbol)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INTERNAL_ERROR, nil)
	}

	appG.OK("取消标记成功")
}

func GetConcerned(c *gin.Context)  {
	var (
		appG = Gin{C: c}
	)

	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _  := strconv.Atoi(c.Query("pageNum"))

	//获得当前用户token
	user, code, err:= serv.GetUserByToken(*c)
	if err != nil {
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}

	//获得关注的所有公司
	companyList, err := serv.GetConcerned(pageNum, pageSize, user.ID)
	data := make(map[string]interface{})
	data["concern"] = companyList
	data["totalNum"] = len(companyList)

	appG.OK(data)

}

func GetHistory(c *gin.Context)  {

}

