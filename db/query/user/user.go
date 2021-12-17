package user

import (
	"context"

	"github.com/CodeHanHan/ferry-backend/db"
	modelUser "github.com/CodeHanHan/ferry-backend/models/user"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
)

func CreateUserRecord(ctx context.Context, record *modelUser.UserTable) error {
	if err := db.Store.Table(modelUser.UserTableName).Create(record).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	return nil
}

func GetByUserName(ctx context.Context, filter *db.Filter) (*modelUser.UserTable, error) {
	var user modelUser.UserTable
	if err := db.Store.Table(modelUser.UserTableName).Where(filter.Params).Find(&user).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}
	return &user, nil
}

func DeleteSysUser(ctx context.Context, id string) error {
	var user modelUser.UserTable
	if err := db.Store.Table(modelUser.UserTableName).Where("id=?", id).Delete(&user).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}
	return nil
}

func UpdateUserRecord(ctx context.Context, filter1 *db.Filter, filter2 *db.Filter) error {
	if err := db.Store.Table(modelUser.UserTableName).Where(filter1.Params).Updates(filter2.Params).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}
	return nil
}
