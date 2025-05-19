package types

type ManagerLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ManagerLoginResp struct {
	Token string `json:"token"`
}
