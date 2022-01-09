package role

import "github.com/CodeHanHan/ferry-backend/models/role"

type CreateRoleRequest struct {
	RoleName string `json:"role_name" form:"role_name" binding:"required"`
	IsAdmin  *int   `json:"is_admin" form:"is_admin" binding:"required,is_admin"`
	Remark   string `json:"remark" form:"remark"`
}

type CreateRoleResponse struct {
	RoleID   string `json:"role_id"`
	RoleName string `json:"role_name"`
	RoleKey  string `json:"role_key"`
}

type DeleteRoleRequest struct {
	RoleID string `json:"role_id" uri:"role_id" binding:"required"`
}

type DeleteRoleResponse struct {
	Result string `json:"result"`
}

type ListRoleRequest struct {
	Offset *int `json:"offset" form:"offset" binding:"required"`
	Limit  int  `json:"limit" form:"limit" binding:"required"`
}

type ListRoleResponse struct {
	Roles  []*role.Role
	Length int
}

type GetRoleRequest struct {
	RoleID string `json:"role_id" uri:"role_id" binding:"required"`
}

type GetRoleResponse struct {
	Role *role.Role
}

type UpdateRoleRequest role.Role

type UpdateRoleResponse struct {
	Result string `json:"result"`
}
