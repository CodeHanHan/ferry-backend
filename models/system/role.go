package system

import (
	"context"
	"errors"
	"time"

	"github.com/CodeHanHan/ferry-backend/db"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/utils/idutil"
)

const RoleTableName = "role"

var (
	ErrUpdateRoleNameNotAllowed = errors.New("角色名称不允许修改！")
	ErrUpdateRoleKeyNotAllowed  = errors.New("角色标识不允许修改！")
)

type IsAdmin int

const (
	NotAdminRole = 0
	AdminRole    = 1
)

type Role struct {
	RoleID     int        `gorm:"primary_key;AUTO_INCREMENT" json:"role_id"`
	RoleName   string     `gorm:"column:role_name" json:"role_name"`
	Status     string     `gorm:"column:status" json:"status"`
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
	Params     string     `json:"params" gorm:"-"`
	MenuIds    []int      `json:"menuIds" gorm:"-"`
	DeptIds    []int      `json:"deptIds" gorm:"-"`
}

func NewRole(rolename string, remark string, admin IsAdmin, createBy string) *Role {
	return &Role{
		RoleName: rolename,
		RoleKey:  idutil.NewHexId(),
		CreateBy: createBy,
		Remark:   remark,
		Admin:    int(admin),
	}
}

// GetPage get role list page by (role_id, role_name, status, role_key), return list, list length and an error
func (r *Role) GetPage(ctx context.Context, limit, offset int) (roles []*Role, count int64, err error) {
	table := db.Store.Select("*").Table("role")
	if r.RoleID != 0 {
		table = table.Where("role_id = ?", r.RoleID)
	}
	if r.RoleName != "" {
		table = table.Where("role_name like ?", "%"+r.RoleName+"%")
	}
	if r.Status != "" {
		table = table.Where("status = ?", r.Status)
	}
	if r.RoleKey != "" {
		table = table.Where("role_key like ?", "%"+r.RoleKey+"%")
	}
	if err := table.Order("role_sort").Offset((offset - 1) * limit).Limit(limit).Find(&roles).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, 0, err
	}
	table.Where("`delete_time` IS NULL").Count(&count)
	return
}

// Get get a role by (role_id, role_name)
func (r *Role) Get(ctx context.Context) (role *Role, err error) {
	table := db.Store.Table(RoleTableName)
	if r.RoleID != 0 {
		table = table.Where("role_id = ?", r.RoleID)
	}
	if role.RoleName != "" {
		table = table.Where("role_name = ?", role.RoleName)
	}
	if err = table.First(role).Error; err != nil {
		logger.Error(ctx, err.Error())
		return
	}

	return
}

// GetList get role list by (role_id, role_name)
func (r *Role) GetList(ctx context.Context) (roles []*Role, err error) {
	table := db.Store.Table(RoleTableName)
	if r.RoleID != 0 {
		table = table.Where("role_id = ?", r.RoleID)
	}
	if r.RoleName != "" {
		table = table.Where("role_name = ?", r.RoleName)
	}
	if err = table.Order("role_sort").Find(&roles).Error; err != nil {
		logger.Error(ctx, err.Error())
		return
	}

	return
}

// Insert create a new role and return its id
func (r *Role) Insert(ctx context.Context) (id int, err error) {
	var i int64
	db.Store.Table("sys_role").Where("(role_name = ? or role_key = ?) and `delete_time` IS NULL", r.RoleName, r.RoleKey).Count(&i)
	if i > 0 {
		return 0, db.ErrDuplicateValue
	}

	r.UpdateBy = ""
	if err := db.Store.Table(RoleTableName).Create(&r).Error; err != nil {
		logger.Error(ctx, err.Error())
		return -1, err
	}

	id = r.RoleID
	return id, nil
}

// Update update a role by id
func (r *Role) Update(ctx context.Context, id int) (updated *Role, err error) {
	if err = db.Store.Table(RoleTableName).First(&updated, id).Error; err != nil {
		logger.Error(ctx, err.Error()) // TODO: not found
		return
	}

	if r.RoleName != "" && r.RoleName != updated.RoleName {
		return updated, ErrUpdateRoleNameNotAllowed
	}

	if r.RoleKey != "" && r.RoleKey != updated.RoleKey {
		return updated, ErrUpdateRoleKeyNotAllowed
	}

	if err := db.Store.Table(RoleTableName).Model(&updated).Updates(&r).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}

	return
}

// BatchDelete batch delete role by given id list and return result
func BatchDelete(ctx context.Context, ids []int) (ok bool, err error) {
	if err := db.Store.Table(RoleTableName).Where("role_id in (?)", ids).Delete(&Role{}).Error; err != nil {
		logger.Error(ctx, err.Error())
		return false, err
	}

	return true, nil
}
