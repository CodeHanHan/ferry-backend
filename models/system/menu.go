package system

import (
	"context"
	"errors"
	"time"

	"github.com/CodeHanHan/ferry-backend/db"
	"github.com/CodeHanHan/ferry-backend/pkg/constants"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/utils/stringutil"
)

const MenuTableName = "menu"

type Menu struct {
	MenuID     int        `gorm:"column:menu_id;primary_key;AUTO_INCREMENT" json:"menu_id"`
	MenuName   string     `gorm:"column:menu_name" json:"menu_name"`
	Title      string     `gorm:"column:title" json:"title"`
	Icon       string     `gorm:"column:icon" json:"icon"`
	Path       string     `gorm:"column:path" json:"path"`           // vue router
	Paths      string     `gorm:"column:paths" json:"paths"`         // router from parent to current node
	MenuType   string     `gorm:"column:menu_type" json:"menu_type"` // C: 菜单， M: 目录， F: 按钮， A: 接口
	Action     string     `gorm:"column:action" json:"action"`
	Permission string     `gorm:"column:permission" json:"permission"` // vue permission
	ParentID   int        `gorm:"column:parent_id" json:"parent_id"`
	NoCache    string     `gorm:"column:no_cache" json:"no_cache"`
	Breadcrumb string     `gorm:"column:breadcrumb" json:"breadcrumb"` // FIXME 未知
	Component  string     `gorm:"column:component" json:"component"`   // 组件路径
	Sort       int        `gorm:"column:sort" json:"sort"`
	Visible    string     `gorm:"column:visible" json:"visible"` // 可见性
	CreateBy   string     `gorm:"column:create_by" json:"create_by"`
	UpdateBy   string     `gorm:"column:update_by" json:"update_by"` // 是否外链
	IsFrame    int        `gorm:"column:is_frame;default:0" json:"is_frame"`
	CreateTime *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime *time.Time `gorm:"column:update_time" json:"update_time"`
	DeleteTime *time.Time `gorm:"column:delete_time" json:"delete_time"`
	Children   []*Menu    `json:"children"`
}

type MenuLabel struct {
	Id       int          `json:"id" gorm:"-"`
	Label    string       `json:"label" gorm:"-"`
	Children []*MenuLabel `json:"children" gorm:"-"`
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

// DeleteMenu delete menu by given menu_id
func (m *Menu) DeleteMenu(ctx context.Context, id int) error {
	if err := db.Store.Table(MenuTableName).Delete(&Menu{}).Where("menu_id = ?", id).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	// 当表行数很多时，考虑使用标记删除，再定时做清理
	return nil
}

// GetMenuById get menu by given menu_id
func (m *Menu) GetMenuById(ctx context.Context, id int) (menu *Menu, err error) {
	if err := db.Store.Table(MenuTableName).First(&menu, id).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}

	return
}

// RecursiveMenu find child menu from params menulist for params menu
func RecursiveMenu(ctx context.Context, menulist []*Menu, menu *Menu) *Menu {
	min := make([]*Menu, 0)

	for i := 0; i < len(menulist); i++ {
		if menu.MenuID != menulist[i].ParentID {
			continue
		}

		mi := Menu{}
		mi.MenuID = menulist[i].MenuID
		mi.MenuName = menulist[i].MenuName
		mi.Title = menulist[i].Title
		mi.Icon = menulist[i].Icon
		mi.Path = menulist[i].Path
		mi.MenuType = menulist[i].MenuType
		mi.Action = menulist[i].Action
		mi.Permission = menulist[i].Permission
		mi.ParentID = menulist[i].ParentID
		mi.NoCache = menulist[i].NoCache
		mi.Breadcrumb = menulist[i].Breadcrumb
		mi.Component = menulist[i].Component
		mi.Sort = menulist[i].Sort
		mi.Visible = menulist[i].Visible
		mi.Children = []*Menu{}

		if mi.MenuType != constants.MenuType_Button {
			ms := RecursiveMenu(ctx, menulist, &mi)
			min = append(min, ms)
		} else {
			min = append(min, &mi)
		}
	}

	menu.Children = min
	return menu
}

// GetPages find menu list by (menu_name, title, visible, menu_type)
func (m *Menu) GetPage(ctx context.Context) (menus []*Menu, err error) {
	table := db.Store.Table(MenuTableName)
	if m.MenuName != "" {
		table = table.Where("menu_name = ?", m.MenuName)
	}
	if m.Title != "" {
		table = table.Where("title like ?", "%"+m.Title+"%")
	}
	if m.Visible != "" {
		table = table.Where("visible = ?", m.Visible)
	}
	if m.MenuType != "" {
		table = table.Where("menu_type = ?", m.MenuType)
	}

	if err = table.Order("sort").Find(&menus).Error; err != nil {
		logger.Error(ctx, err.Error())
		return
	}
	return
}

// GetPages find menu list by (menu_name, path, action, menu_type)
func (e *Menu) Get(ctx context.Context) (menus []*Menu, err error) {
	table := db.Store.Table(MenuTableName)
	if e.MenuName != "" {
		table = table.Where("menu_name = ?", e.MenuName)
	}
	if e.Path != "" {
		table = table.Where("path = ?", e.Path)
	}
	if e.Action != "" {
		table = table.Where("action = ?", e.Action)
	}
	if e.MenuType != "" {
		table = table.Where("menu_type = ?", e.MenuType)
	}

	if err = table.Order("sort").Find(&menus).Error; err != nil {
		logger.Error(ctx, err.Error())
		return
	}
	return
}

// SetMenu find menu list by given condition and set children menu for each menu
func (m *Menu) SetMenu(ctx context.Context) (menus []*Menu, err error) {
	menulist, err := m.GetPage(ctx)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(menulist); i++ {
		if menulist[i].ParentID != 0 {
			continue
		}
		menu_ := RecursiveMenu(ctx, menulist, menulist[i])

		menus = append(menus, menu_)
	}
	return
}

// SetMenuLabel set menulabel list by give menu
func (m *Menu) SetMenuLabel(ctx context.Context) (mls []*MenuLabel, err error) {
	menulist, err := m.Get(ctx)
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}

	mls = make([]*MenuLabel, 0)
	for i := 0; i < len(menulist); i++ {
		if menulist[i].ParentID != 0 {
			continue
		}
		ml := MenuLabel{}
		ml.Id = menulist[i].MenuID
		ml.Label = menulist[i].Title

		menus := RecursiveMenuLabel(ctx, menulist, &ml)

		mls = append(mls, menus)
	}
	return
}

// RecursiveMenuLabel find router tree for given menulabel
func RecursiveMenuLabel(ctx context.Context, menulist []*Menu, menu *MenuLabel) *MenuLabel {
	min := make([]*MenuLabel, 0)
	for j := 0; j < len(menulist); j++ {
		if menu.Id != menulist[j].ParentID {
			continue
		}

		mi := MenuLabel{}
		mi.Id = menulist[j].MenuID
		mi.Label = menulist[j].Title
		mi.Children = make([]*MenuLabel, 0)
		if menulist[j].MenuType != constants.MenuType_Button {
			ms := RecursiveMenuLabel(ctx, menulist, &mi)
			min = append(min, ms)
		} else {
			min = append(min, &mi)
		}
	}
	menu.Children = min
	return menu
}
