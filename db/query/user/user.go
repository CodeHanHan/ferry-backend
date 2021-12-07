package user

import (
	"context"

	"github.com/CodeHanHan/ferry-backend/db"
	users "github.com/CodeHanHan/ferry-backend/models/users"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
)

func CreateUserRecord(ctx context.Context, record *users.UsersTable) error {
	if err := db.Store.Table(users.UsersTableName).Create(record).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	return nil
}
