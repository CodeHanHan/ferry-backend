package ping

import (
	"time"

	"github.com/CodeHanHan/ferry-backend/utils/idutil"
)

const PingRecordTableName = "ping_record"

type PingRecord struct {
	PingID     string    `gorm:"column:ping_id;primary_key"`
	Message    string    `gorm:"column:message"`
	Reply      string    `gorm:"column:reply"`
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`
}

func NewPingRecord(message string, reply string) *PingRecord {
	return &PingRecord{
		PingID:  idutil.NewHexId(),
		Message: message,
		Reply:   reply,
	}
}
