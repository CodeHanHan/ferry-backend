package post

import (
	"time"

	"github.com/CodeHanHan/ferry-backend/utils/idutil"
)

const PostTableName = "post"

type PostStatus int

const (
	PostDeactivated PostStatus = 0
	PostActive      PostStatus = 1
)

type Post struct {
	PostID     string     `gorm:"column:post_id;primary_key" json:"post_id"`
	PostName   string     `gorm:"column:post_name" json:"post_name"`
	PostCode   string     `gorm:"column:post_code" json:"post_code"`
	Sort       int        `gorm:"column:sort" json:"sort"`
	Status     int        `gorm:"column:status" json:"status"`
	Remark     string     `gorm:"column:remark" json:"remark"`
	CreateBy   string     `gorm:"column:create_by" json:"create_by"`
	UpdateBy   string     `gorm:"column:update_by" json:"update_by"`
	CreateTime *time.Time `gorm:"column:create_time" json:"create_time" default:"2000-01-01 00:00:00"`
	UpdateTime *time.Time `gorm:"column:update_time" json:"update_time" default:"2000-01-01 00:00:00"`
	DeleteTime *time.Time `gorm:"column:delete_time" json:"delete_time" default:"2000-01-01 00:00:00"`
}

func NewPost(postname string, postcode string, create_by string) *Post {
	return &Post{
		PostID:   idutil.GetId("post"),
		PostName: postname,
		PostCode: postcode,
		Status:   int(PostActive),
		CreateBy: create_by,
	}
}
