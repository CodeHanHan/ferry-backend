package form

type PingRequest struct {
	Message string `form:"message" binding:"required"`
}

type ListPingRequest struct {
	Offset int `form:"offset" binding:"gte=-1"`
	Limit  int `form:"limit" binding:"required,gte=0"`
}
