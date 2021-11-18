package ping

import (
	"github.com/CodeHanHan/ferry-backend/db"
	"github.com/CodeHanHan/ferry-backend/models/ping"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
)

func CreatePingRecord(record *ping.PingRecord) error {
	if err := db.Store.Table(ping.PingRecordTableName).Create(record).Error; err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
