package post

// import (
// 	"github.com/CodeHanHan/ferry-backend/models/post"
// )

// type CreatePostRequest struct {
// 	PostCode string `json:"post_code" form:"post_code" binding:"required"`
// 	PostName string `json:"post_name" form:"post_name" binding:"required"`
// }

// type CreatePostResponse struct {
// 	PostID   string `json:"post_id"`
// 	PostName string `json:"post_name"`
// }

// type DeletePostRequest struct {
// 	PostID string `json:"post_id" uri:"post_id" binding:"required"`
// }

// type DeletePostResponse struct {
// 	Result string `json:"result"`
// }

// type ListPostRequest struct {
// 	Offset *int `json:"offset" form:"offset" binding:"required"`
// 	Limit  int  `json:"limit" form:"limit" binding:"required"`
// }

// type ListPostResponse struct {
// 	Post   []*post.Post
// 	Length int
// }

// type GetPostRequest struct {
// 	PostID string `json:"post_id" uri:"post_id" binding:"required"`
// }

// type GetPostResponse struct {
// 	Post *post.Post
// }

// type UpdatePostRequest post.Post

// type UpdatePostResponse struct {
// 	Result string `json:"result"`
// }
