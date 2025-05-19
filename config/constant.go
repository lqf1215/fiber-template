package config

const (
	LOCAL_USERID_UINT  = "user_id_uint"
	LOCAL_USERID_INT64 = "user_id_int64"
	LOCAL_TOKEN        = "token"

	MANAGER_LOCAL_USERID_UINT  = "manager_user_id_uint"
	MANAGER_LOCAL_USERID_INT64 = "manager_user_id_int64"
	MANAGER_LOCAL_USER_NAME    = "manager_user_name" //管理员用户名
	MANAGER_LOCAL_USER_ROLE    = "manager_user_role" //管理系统角色
	MANAGER_LOCAL_USER         = "manager__user"     //管理系统用户
)

const MESSAGE_SUCCESS = 0
const MESSAGE_FAIL = -1
const TOKEN_FAIL = -2
const OPERATION_FAIL = -3

const LOGIN_FILE_KEY = "panasia-sit@2020"

const LOGIN_P12_PATH = "file/10000006.p12"

const LOGIN_CRT_PATH = "file/10000006.crt"
