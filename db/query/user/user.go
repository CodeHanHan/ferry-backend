package user

import (
	"context"

	"github.com/CodeHanHan/ferry-backend/db"
	"github.com/CodeHanHan/ferry-backend/models/ping"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
)

func CreateUserRecord(ctx context.Context, record *ping.UsersTable) error {
	if err := db.Store.Table(ping.UsersTableName).Create(record).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	return nil
}
