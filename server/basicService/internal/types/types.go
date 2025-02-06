// Code generated by goctl. DO NOT EDIT.
package types

type GetUserInfoReq struct {
	Id int64 `json:"id"`
}

type GetUserInfoResp struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Token    string `json:"token"`
	ExpireAt string `json:"expireAt"`
}

type UpdateUserInfoReq struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}
