package menurole

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
