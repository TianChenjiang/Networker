package e

const (
	SUCCESS                        = 200
	BAD_REQUEST                    = 400
	UNAUTHENTIC                    = 401
	METHOD_FORBIDDEN               = 403
	RESOURCE_NOT_FOUND             = 404

	INTERNAL_ERROR                 = 500
	BAD_GATEWAY                    = 502
	SERVICE_UNAVAILABLE            = 503
	GATEWAY_TIMEOUT                = 504

	ERROR_GET_USERLIST             = 10001
	ERROR_GET_USER                 = 10002
	ERROR_CREATE_USER              = 10003
	ERROR_EDIT_USER                = 10004
	ERROR_DELETE_USER              = 10005
	ERROR_CHANGE_PASSWORD_SAME     = 10006
	ERROR_CHANGE_PASSWORD          = 10007


	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

)
