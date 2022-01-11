package role

import (
	"time"

	"github.com/CodeHanHan/ferry-backend/utils/idutil"
)

const RoleTableName = "role"

type IsAdmin int

const (
	NotAdminRole = 0
	AdminRole    = 1
)

type RoleStatus int

const (
	RoleDeactivated RoleStatus = 0
	RoleActive      RoleStatus = 1
)

type Role struct {
	RoleID     string     `gorm:"column:role_id;primary_key;" json:"role_id"`
	RoleName   string     `gorm:"column:role_name" json:"role_name"`
	Status     int        `gorm:"column:status" json:"status"`
	RoleKey    string     `gorm:"column:role_key" json:"role_key"`
	RoleSort   int        `gorm:"column:role_sort" json:"role_sort"`
	Flag       string     `gorm:"column:flag" json:"flag"`
	CreateBy   string     `gorm:"column:create_by" json:"create_by"`
	UpdateBy   string     `gorm:"column:update_by" json:"update_by"`
	Remark     string     `gorm:"column:remark" json:"remark"`
	Admin      int        `gorm:"column:admin" json:"admin"`
	CreateTime *time.Time `gorm:"column:create_time" json:"create_time" default:"2000-01-01 00:00:00"`
	UpdateTime *time.Time `gorm:"column:update_time" json:"update_time" default:"2000-01-01 00:00:00"`
	DeleteTime *time.Time `gorm:"column:delete_time" json:"delete_time" default:"2000-01-01 00:00:00"`
}

func NewRole(rolename string, remark string, admin IsAdmin, createBy string) *Role {
	return &Role{
		RoleID:   idutil.GetId("role"),
		RoleName: rolename,
		Status:   int(RoleActive),
		RoleKey:  idutil.NewHexId(),
		CreateBy: createBy,
		Remark:   remark,
		Admin:    int(admin),
	}
}
