package system

import (
	"context"
	"time"

	"github.com/CodeHanHan/ferry-backend/db"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
)

const PostTableName = "post"

type PostStatus int

type Post struct {
	PostId     int        `gorm:"primary_key;AUTO_INCREMENT" json:"postId"` //岗位编号
	PostName   string     `gorm:"type:varchar(128);" json:"postName"`       //岗位名称
	PostCode   string     `gorm:"type:varchar(128);" json:"postCode"`       //岗位代码
	Sort       int        `gorm:"type:int(4);" json:"sort"`                 //岗位排序
	Status     string     `gorm:"type:int(1);" json:"status"`               //状态
	Remark     string     `gorm:"type:varchar(255);" json:"remark"`         //描述
	CreateBy   string     `gorm:"type:varchar(128);" json:"createBy"`
	UpdateBy   string     `gorm:"type:varchar(128);" json:"updateBy"`
	Params     string     `gorm:"-" json:"params"`
	CreateTime *time.Time `gorm:"column:create_time" json:"create_time" default:"2000-01-01 00:00:00"`
	UpdateTime *time.Time `gorm:"column:update_time" json:"update_time" default:"2000-01-01 00:00:00"`
	DeleteTime *time.Time `gorm:"column:delete_time" json:"delete_time" default:"2000-01-01 00:00:00"`
}

func (p *Post) Create(ctx context.Context) (*Post, error) {
	var post Post
	if err := db.Store.Table(PostTableName).Create(p).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}

	post = *p
	return &post, nil
}
