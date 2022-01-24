package system

import (
	"context"

	"github.com/CodeHanHan/ferry-backend/db"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/utils/stringutil"
)

const (
	BindRoleMenuTableName = "bind_role_menu"
)

type BindRoleMenu struct {
	RoleID   int    `gorm:"column:role_id" json:"role_id"`
	MenuID   int    `gorm:"column:menu_id" json:"menu_id"`
	RoleName string `gorm:"column:role_name" json:"role_name"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	ID       int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
}

// Get find bindRoleMenu models by give id, if id == 0, then find all
func (rm *BindRoleMenu) Get(ctx context.Context) (rms []*BindRoleMenu, err error) {
	table := db.Store.Table(BindRoleMenuTableName)
	if rm.RoleID != 0 { // 如果设置为0，代表查找全部
		table = table.Where("role_id = ?", rm.RoleID)
	}

	if err := table.Find(&rms).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}

	return
}

func (rm *BindRoleMenu) Delete(ctx context.Context, roleId, menuId string) (err error) {
	rm.RoleID, err = stringutil.String2Int(roleId)
	if err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	table := db.Store.Table(BindRoleMenuTableName).Where("role_id = ?", roleId)
	if menuId != "" {
		table = table.Where("menu_id  = ?", menuId)
	}

	if err := table.Delete(&rm).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	return nil
}
