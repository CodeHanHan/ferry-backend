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

func ListPingRecords(ctx context.Context, offset, limit int) ([]*ping.PingRecord, error) {
	var ans []*ping.PingRecord = make([]*ping.PingRecord, 0)
	if err := db.Store.Table(ping.PingRecordTableName).Limit(limit).Offset(offset).Find(&ans).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}
	return ans, nil
}

func DeletePingRecord(ctx context.Context, pk string) error {
	if err := db.Store.Table(ping.PingRecordTableName).Where(map[string]interface{}{"ping_id": pk}).Delete(&ping.PingRecord{}).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	return nil
}

func GetPingRecordByPingID(ctx context.Context, pk string) (record *ping.PingRecord, err error) {
	if err := db.Store.Table(ping.PingRecordTableName).Where(map[string]interface{}{"ping_id": pk}).Take(&record).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}

	return
}

func UpdatePingRecord(ctx context.Context, record *ping.PingRecord) error {
	if err := db.Store.Table(ping.PingRecordTableName).Updates(record).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	return nil
}
