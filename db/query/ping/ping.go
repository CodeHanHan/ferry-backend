package ping

import (
	"context"

	"github.com/CodeHanHan/ferry-backend/db"
	ping "github.com/CodeHanHan/ferry-backend/models/ping"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
)

func CreatePingRecord(ctx context.Context, record *ping.PingRecord) error {
	if err := db.Store.Table(ping.PingRecordTableName).Create(record).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	return nil
}

func PagePingRecords(ctx context.Context, offset, limit int) ([]*ping.PingRecord, error) {
	var ans []*ping.PingRecord = make([]*ping.PingRecord, 0)
	if err := db.Store.Table(ping.PingRecordTableName).Limit(limit).Offset(offset).Find(&ans).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}
	return ans, nil
}

func DeletePingRecord(ctx context.Context, pk string) error {
	if err := db.Store.Table(ping.PingRecordTableName).Where("ping_id = ?", pk).
		Delete(&ping.PingRecord{}).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	return nil
}

func UpdatePingRecord(ctx context.Context, pk string, message string, reply string) error {
	res := db.Store.Table(ping.PingRecordTableName).Where("ping_id = ?", pk).Updates(map[string]interface{}{"message": message, "reply": reply})
	if res.Error != nil {
		logger.Error(ctx, res.Error.Error())
		return res.Error
	}
	if res.RowsAffected <= 0 {
		logger.Error(ctx, "要更新的值不存在")
		return db.ErrNotExist
	}

	return nil
}
