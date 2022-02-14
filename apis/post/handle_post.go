package post

// import (
// 	"errors"
// 	"time"

// 	"github.com/CodeHanHan/ferry-backend/db"
// 	post "github.com/CodeHanHan/ferry-backend/db/query/post"
// 	modelPost "github.com/CodeHanHan/ferry-backend/models/post"
// 	"github.com/CodeHanHan/ferry-backend/pkg/app"
// 	formPost "github.com/CodeHanHan/ferry-backend/pkg/form/post"
// 	"github.com/CodeHanHan/ferry-backend/pkg/logger"
// 	"github.com/CodeHanHan/ferry-backend/pkg/sender"
// 	"github.com/gin-gonic/gin"
// 	"github.com/gin-gonic/gin/binding"
// 	"gorm.io/gorm"
// )

// // CreatePost godoc
// // @Summary 创建岗位
// // @Description 根据PostName和PostCode创建岗位信息
// // @Tags post
// // @ID post-create
// // @Param post body formPost.CreatePostRequest true "post名称和post等级"
// // @Success 200 {object} formPost.CreatePostResponse
// // @Failure 500 {object} app.ErrResponse
// // @Failure 400 {object} app.ErrResponse
// // @Produce  json
// // @Router /post [post]
// // @Security BearerAuth
// func CreatePost(c *gin.Context) {
// 	var req formPost.CreatePostRequest
// 	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
// 		logger.ErrorParams(c, err)
// 		app.ErrorParams(c, err)
// 		return
// 	}

// 	creator, _, err := sender.GetSender(c)
// 	if err != nil {
// 		logger.Error(c, err.Error())
// 		app.InternalServerError(c)
// 		return
// 	}

// 	newPost := modelPost.NewPost(req.PostName, req.PostCode, creator)
// 	if err := post.CreatePost(c, newPost); err != nil {
// 		if errors.Is(db.ErrDuplicateValue, err) {
// 			app.Error(c, app.Err_Invalid_Argument, "PostName already exists")
// 			return
// 		}
// 		app.InternalServerError(c)
// 		return
// 	}

// 	resp := formPost.CreatePostResponse{
// 		PostID:   newPost.PostID,
// 		PostName: newPost.PostName,
// 	}

// 	app.OK(c, resp)
// }

// // DeletePost godoc
// // @Summary 删除岗位
// // @Description 根据PostID删除岗位
// // @Tags post
// // @ID post-delete
// // @Param post_id path string true "岗位唯一id"
// // @Success 200 {object} formPost.DeletePostResponse
// // @Failure 500 {object} app.ErrResponse
// // @Failure 400 {object} app.ErrResponse
// // @Produce  json
// // @Router /post/{post_id} [delete]
// // @Security BearerAuth
// func DeletePost(c *gin.Context) {
// 	var req formPost.DeletePostRequest
// 	if err := c.ShouldBindUri(&req); err != nil {
// 		logger.ErrorParams(c, err)
// 		app.ErrorParams(c, err)
// 		return
// 	}

// 	if err := post.DeletePostById(c, req.PostID); err != nil {
// 		app.InternalServerError(c)
// 		return
// 	}

// 	app.OK(c, formPost.DeletePostResponse{
// 		Result: "success",
// 	})
// }

// // ListPost godoc
// // @Summary 查询岗位列表
// // @Description 根据offset和limit查询岗位列表
// // @Tags post
// // @ID post-list
// // @Param offset query int true "偏移"
// // @Param limit query int true "限制"
// // @Success 200 {object} formPost.ListPostResponse
// // @Failure 500 {object} app.ErrResponse
// // @Failure 400 {object} app.ErrResponse
// // @Produce  json
// // @Router /post [get]
// // @Security BearerAuth
// func ListPost(c *gin.Context) {
// 	var req formPost.ListPostRequest
// 	if err := c.ShouldBind(&req); err != nil {
// 		logger.ErrorParams(c, err)
// 		app.ErrorParams(c, err)
// 		return
// 	}

// 	list, err := post.SearchPost(c, *req.Offset, req.Limit)
// 	if err != nil {
// 		app.InternalServerError(c)
// 		return
// 	}

// 	resp := formPost.ListPostResponse{
// 		Post:   list,
// 		Length: len(list),
// 	}

// 	app.OK(c, resp)
// }

// // GetPost godoc
// // @Summary 查询岗位
// // @Description 根据岗位id查询岗位信息
// // @Tags post
// // @ID post-get
// // @Param post_id path string true "部门id"
// // @Success 200 {object} formPost.GetPostResponse
// // @Failure 500 {object} app.ErrResponse
// // @Failure 400 {object} app.ErrResponse
// // @Produce  json
// // @Router /post/{post_id} [get]
// // @Security BearerAuth
// func GetPost(c *gin.Context) {
// 	var req formPost.GetPostRequest
// 	if err := c.ShouldBindUri(&req); err != nil {
// 		logger.ErrorParams(c, err)
// 		app.ErrorParams(c, err)
// 		return
// 	}

// 	f := db.NewFilter().Set("post_id", req.PostID)
// 	post, err := post.GetPost(c, f)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			app.Errorf(c, app.Err_Not_found, "查询失败，未找到该记录值: %s", req.PostID)
// 			return
// 		}
// 		app.InternalServerError(c)
// 		return
// 	}

// 	resp := formPost.GetPostResponse{
// 		Post: post,
// 	}

// 	app.OK(c, resp)
// }

// // UpdatePost godoc
// // @Summary 更新岗位
// // @Description 更新岗位信息
// // @Tags post
// // @ID post-update
// // @Param post body formPost.UpdatePostRequest true "包含postid、post名称、post等级等相关信息"
// // @Success 200 {object} formPost.UpdatePostResponse
// // @Failure 500 {object} app.ErrResponse
// // @Failure 400 {object} app.ErrResponse
// // @Accept application/json
// // @Produce  json
// // @Router /post [put]
// // @Security BearerAuth
// func UpdatePost(c *gin.Context) {
// 	var req formPost.UpdatePostRequest
// 	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
// 		logger.ErrorParams(c, err)
// 		app.ErrorParams(c, err)
// 		return
// 	}

// 	post_ := modelPost.Post(req)

// 	f := db.NewFilter().Set("post_id", req.PostID)
// 	_, err := post.GetPost(c, f)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			app.Errorf(c, app.Err_Not_found, "查询失败，未找到该记录值: %s", req.PostID)
// 			return
// 		}
// 		app.InternalServerError(c)
// 		return
// 	}

// 	sender, _, err := sender.GetSender(c)
// 	if err != nil {
// 		logger.Error(c, err.Error())
// 		app.Error(c, app.Err_Unauthenticated, err)
// 		return
// 	}

// 	var TimeNow = time.Now()
// 	post_.UpdateBy = sender
// 	post_.UpdateTime = &TimeNow
// 	if err := post.UpdatePost(c, &post_); err != nil {
// 		app.InternalServerError(c)
// 		return
// 	}

// 	resp := formPost.UpdatePostResponse{
// 		Result: "success",
// 	}

// 	app.OK(c, resp)
// }
