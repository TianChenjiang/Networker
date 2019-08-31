package e

var msgFlags = map[int]string{
	SUCCESS:                         "ok",
	BAD_REQUEST:                     "请求参数错误",
	METHOD_FORBIDDEN:                "服务器拒绝请求",
	UNAUTHENTIC:                     "未授权",
	//RESOURCE_NOT_FOUND:              "请求的URL不存在",
	INTERNAL_ERROR:                  "内部错误",
	//BAD_GATEWAY:                     "网关错误",
	//SERVICE_UNAVAILABLE:             "服务不可用",
	//GATEWAY_TIMEOUT:                 "网关超时",
	ERROR_GET_USERLIST:              "获得用户列表出错",
	ERROR_GET_USER:                  "获得用户出错",
	ERROR_CREATE_USER:               "注册用户参数错误",
	ERROR_DELETE_USER:               "删除用户失败",
	ERROR_CHANGE_PASSWORD_SAME:      "密码与原来一致",
	ERROR_CHANGE_PASSWORD:           "修改密码错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存上传图片失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "密码错误或者token错误",

}

func GetMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}
	return msgFlags[500]
}
