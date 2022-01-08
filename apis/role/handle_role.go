package role

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/CodeHanHan/ferry-backend/db"
	role "github.com/CodeHanHan/ferry-backend/db/query/role"
	modelRole "github.com/CodeHanHan/ferry-backend/models/role"
	"github.com/CodeHanHan/ferry-backend/pkg/app"
	formRole "github.com/CodeHanHan/ferry-backend/pkg/form/role"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/pkg/sender"
)

// CreateRole godoc
// @Summary 创建角色
// @Description 根据一个RoleName创建角色，RoleName 要求不能重复
// @Tags role
// @ID role-create
// @Param role body formRole.CreateRoleRequest true "remark 可空"
// @Success 200 {object} formRole.CreateRoleResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Accept application/json
// @Produce  json
// @Router /role [post]
// @Security BearerAuth
func CreateRole(c *gin.Context) {
	var req formRole.CreateRoleRequest
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		logger.ErrorParams(c, err)
		app.ErrorParams(c, err)
		return
	}

	creator, _, err := sender.GetSender(c)
	if err != nil {
		logger.Error(c, err.Error())
		app.InternalServerError(c)
		return
	}

	newRole := modelRole.NewRole(req.RoleName, req.Remark, modelRole.IsAdmin(*req.IsAdmin), creator)
	if err := role.CreateRole(c, newRole); err != nil {
		if errors.Is(db.ErrDuplicateValue, err) {
			app.Error(c, app.Err_Invalid_Argument, "RoleName already exists")
			return
		}
		app.InternalServerError(c)
		return
	}

	resp := formRole.CreateRoleResponse{
		RoleID:   newRole.RoleID,
		RoleName: newRole.RoleName,
		RoleKey:  newRole.RoleKey,
	}

	app.OK(c, resp)
}

// DeleteRole godoc
// @Summary 删除角色
// @Description 删除角色，幂等操作
// @Tags role
// @ID role-delete
// @Param role_id path string true "角色唯一id"
// @Success 200 {object} formRole.DeleteRoleResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Accept application/json
// @Produce  json
// @Router /role/{role_id} [delete]
// @Security BearerAuth
func DeleteRole(c *gin.Context) {
	var role_ formRole.DeleteRoleRequest
	if err := c.ShouldBindUri(&role_); err != nil {
		logger.ErrorParams(c, err)
		app.ErrorParams(c, err)
		return
	}

	if err := role.DeleteRoleById(c, role_.RoleID); err != nil {
		app.InternalServerError(c)
		return
	}

	app.OK(c, formRole.DeleteRoleResponse{
		Result: "success",
	})
}

// DeleteRole godoc
// @Summary 查询角色列表
// @Description 根据offset和limit查询角色列表
// @Tags role
// @ID role-list
// @Param offset query int true "偏移"
// @Param limit query int true "限制"
// @Success 200 {object} formRole.ListRoleResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Produce  json
// @Router /role [get]
// @Security BearerAuth
func ListRoles(c *gin.Context) {
	var req formRole.ListRoleRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.ErrorParams(c, err)
		app.ErrorParams(c, err)
		return
	}

	list, err := role.SearchRole(c, *req.Offset, req.Limit)
	if err != nil {
		app.InternalServerError(c)
		return
	}

	resp := formRole.ListRoleResponse{
		Roles:  list,
		Length: len(list),
	}

	app.OK(c, resp)
}
