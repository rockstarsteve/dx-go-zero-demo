// Code generated by goctl. DO NOT EDIT.
package types

type OrderReq struct {
	Id string `path:"id"`
}

type OrderReply struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Age    int64  `json:"age"`
	Gender string `json:"gender"`
}

type SaveReq struct {
}

type SaveResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}
