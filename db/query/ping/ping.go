package ping

import (
	"context"

	"github.com/CodeHanHan/ferry-backend/db"
	"github.com/CodeHanHan/ferry-backend/models/ping"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
)

func CreatePingRecord(ctx context.Context, record *ping.PingRecord) error {
	if err := db.Store.Table(ping.PingRecordTableName).Create(record).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	return nil
}
