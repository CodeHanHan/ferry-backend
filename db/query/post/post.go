package post

import (
	"context"

	"github.com/CodeHanHan/ferry-backend/db"
	modelPost "github.com/CodeHanHan/ferry-backend/models/post"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/go-sql-driver/mysql"
)

func CreatePost(ctx context.Context, post *modelPost.Post) error {
	if err := db.Store.Table(modelPost.PostTableName).Create(post).Error; err != nil {
		logger.Error(ctx, err.Error())
		err, ok := err.(*mysql.MySQLError)
		if ok && err.Number == 1062 {
			return db.ErrDuplicateValue
		}

		return err
	}
	return nil
}

func DeletePostById(ctx context.Context, postId string) error {
	var post modelPost.Post
	if err := db.Store.Table(modelPost.PostTableName).Where(map[string]interface{}{"post_id": postId}).Delete(&post).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}
	return nil
}

func SearchPost(ctx context.Context, offset, limit int) (list []*modelPost.Post, err error) {
	if err := db.Store.Table(modelPost.PostTableName).Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}

	return
}

func GetPost(ctx context.Context, f *db.Filter) (post *modelPost.Post, err error) {
	if err := db.Store.Table(modelPost.PostTableName).Where(f.Params).Take(&post).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}

	return
}

func UpdatePost(ctx context.Context, post *modelPost.Post) error {
	if err := db.Store.Table(modelPost.PostTableName).Updates(post).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	return nil

}