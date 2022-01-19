package app

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func (res *Response) ReturnOK() *Response {
	res.Code = 200
	return res
}
