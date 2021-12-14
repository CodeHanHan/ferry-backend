package form

import (
	"github.com/CodeHanHan/ferry-backend/models/ping"
)

// post: /ping
type PingRequest struct {
	Message string `form:"message" binding:"required"`
}

type PingResponse struct {
	Reply string `json:"reply"`
}

// get: /ping
type ListPingRequest struct {
	Offset int `form:"offset" binding:"required,gte=-1"`
	Limit  int `form:"limit" binding:"required,gte=0"`
}

type ListPingResponse struct {
	Records []*ping.PingRecord `json:"records"`
}

// delete: /ping
type DeletePingRequest struct {
	PingID string `form:"ping_id" binding:"required"`
}

type DeletePingResponse struct {
	Result string `json:"result"`
}

// update: /ping
type UpdatePingRequest struct {
	PingID        string `form:"ping_id" binding:"required"`
	UpdateMessage string `form:"updatemessage" binding:"required"`
}

type UpdatePingResponse struct {
	Result string `json:"result"`
}
