package web

import (
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
	err := c.Bind(&schema)
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
		schema schema.Auth
	)

	err := c.BindJSON(&schema)


	isExist, err := serv.Check(*c, schema.Email, schema.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(schema.Email, schema.Password)
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
	_, image, err := c.Request.FormFile("image")
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INTERNAL_ERROR, nil)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, e.BAD_REQUEST, nil)
		return
	}

	if err := c.SaveUploadedFile(image, "library/pic/"+image.Filename); err != nil { //todo 保存路径
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}
	appG.OK(image.Filename)
	return

}

