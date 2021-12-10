package user

import (
	"context"

	"github.com/CodeHanHan/ferry-backend/db"
	modelUsers "github.com/CodeHanHan/ferry-backend/models/users"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
)

func CreateUserRecord(ctx context.Context, record *modelUsers.UsersTable) error {
	if err := db.Store.Table(modelUsers.UsersTableName).Create(record).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	return nil
}

func GetByUserName(ctx context.Context, username string) (*modelUsers.UsersTable, error) {
	var user modelUsers.UsersTable
	if err := db.Store.Table(modelUsers.UsersTableName).Where("user_name = ?", username).Find(&user).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}
	return &user, nil
}
