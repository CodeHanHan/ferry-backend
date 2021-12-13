package app

type ErrorResponse struct {
	Code    ErrCode     `json:"code"`
	Message string      `json:"message"`
	Detail  interface{} `json:"details"`
}
