package api

const (
	Unauthorized = 401
	Forbidden    = 403
	NotFound     = 404
	ServerError  = 500
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
