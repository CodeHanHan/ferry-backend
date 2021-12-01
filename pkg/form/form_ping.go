package form

type PingRequest struct {
	Message string `form:"message" binding:"required"`
}

type ListPingRequest struct {
	Offset int `form:"offset" binding:"required,gte=-1"`
	Limit  int `form:"limit" binding:"required,gte=0"`
}

type DeletePingRequest struct {
	PingID string `form:"ping_id" binding:"required"`
}

type UpdatePingRequest struct {
	PingID        string `form:"ping_id" binding:"required"`
	UpdateMessage string `form:"updatemessage" binding:"required"`
}
