package menu

import (
	"context"
	"errors"
	"time"

	"github.com/CodeHanHan/ferry-backend/db"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/utils/stringutil"
)

const MenuTableName = "menu"

type Menu struct {
	MenuID     int        `gorm:"column:menu_id;primary_key;AUTO_INCREMENT" json:"menu_id"`
	MenuName   string     `gorm:"column:menu_name" json:"menu_name"`
	Title      string     `gorm:"column:title" json:"title"`
	Icon       string     `gorm:"column:icon" json:"icon"`
	Path       string     `gorm:"column:path" json:"path"`
	Paths      string     `gorm:"column:paths" json:"paths"`
	MenuType   string     `gorm:"column:menu_type" json:"menu_type"`
	Action     string     `gorm:"column:action" json:"action"`
	Permission string     `gorm:"column:permission" json:"permission"`
	ParentID   int        `gorm:"column:parent_id" json:"parent_id"`
	NoCache    string     `gorm:"column:no_cache" json:"no_cache"`
	Breadcrumb string     `gorm:"column:breadcrumb" json:"breadcrumb"`
	Component  string     `gorm:"column:component" json:"component"`
	Sort       int        `gorm:"column:sort" json:"sort"`
	Visible    string     `gorm:"column:visible" json:"visible"`
	CreateBy   string     `gorm:"column:create_by" json:"create_by"`
	UpdateBy   string     `gorm:"column:update_by" json:"update_by"`
	IsFrame    int        `gorm:"column:is_frame;default:0" json:"is_frame"`
	CreateTime *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime *time.Time `gorm:"column:update_time" json:"update_time"`
	DeleteTime *time.Time `gorm:"column:delete_time" json:"delete_time"`
}

// Create creates a menu object, return its id and an error if it exists
func (m *Menu) Create(ctx context.Context) (id int, err error) {
	if err := db.Store.Table(MenuTableName).Create(&m).Error; err != nil {
		logger.Error(ctx, err.Error())
		return -1, err
	}

	// InitPaths(ctx, m)

	return m.MenuID, nil
}

// InitPaths init menu path when a menu is created. menu.path = parentMenu.path + "/" + menu.id
func InitPaths(ctx context.Context, m *Menu) error {
	parentMenu := new(Menu)
	if int(m.ParentID) != 0 { // not top menu
		db.Store.Table(MenuTableName).Where("munu_id = ?", m.ParentID).First(parentMenu)
		if parentMenu.Paths == "" {
			return errors.New("父级paths异常，请尝试对当前节点父级菜单进行更新操作!")
		}
		m.Paths = stringutil.Join(parentMenu.Paths, "/", stringutil.Int2String(m.MenuID))
	} else { // top menu
		m.Paths = stringutil.Join("/0/", stringutil.Int2String(m.MenuID))
	}

	if err := db.Store.Table(MenuTableName).Where("menu_id = ?", m.MenuID).Update("paths", m.Paths).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	return nil
}

// Update select a menu by id for update
// TODO: transaction update
func (m *Menu) Update(ctx context.Context, id int) (menu *Menu, err error) {
	if err := db.Store.Table(MenuTableName).First(&menu, id).Error; err != nil {
		logger.Error(ctx, err.Error()) // TODO: not found error
		return nil, err
	}

	if err := db.Store.Table(MenuTableName).Model(&menu).Updates(&m).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}

	if err := InitPaths(ctx, m); err != nil {
		return nil, err
	}

	return
}

func (m *Menu) DeleteMenu(ctx context.Context, id int) error {
	if err := db.Store.Table(MenuTableName).Delete(&Menu{}).Where("menu_id = ?", id).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	// 当表行数很多时，考虑使用标记删除，再定时做清理
	return nil
}

func (m *Menu) GetMenuById(ctx context.Context, id int) (menu *Menu, err error) {
	if err := db.Store.Table(MenuTableName).First(&menu, id).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}

	return
}
