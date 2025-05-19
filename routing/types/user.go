package types

type RegisterReq struct {
	LoginType       string `json:"login_type"` // 登录类型 1:手机号 2:邮箱 默认手机号
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Area            string `json:"area"`
	LoginPwd        string `json:"login_pwd"` // 登录密码
	ConfirmLoginPwd string `json:"confirm_login_pwd"`
}

type UserInfoResp struct {
	AvatarUrl string `json:"avatar_url"` // 头像
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

// LoginReq 登录请求
type LoginReq struct {
	LoginType string `json:"login_type"` // 登录类型 1:手机号 2:邮箱 默认手机号
	Area      string `json:"area"`       // 区号
	Phone     string `json:"phone"`      // 手机号
	Email     string `json:"email"`      // 邮箱
	LoginPwd  string `json:"login_pwd"`  // 登录密码
}

type LoginResp struct {
	Token string `json:"token"` // token
}
