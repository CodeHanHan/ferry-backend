package system

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"

	"github.com/CodeHanHan/ferry-backend/db"
	role "github.com/CodeHanHan/ferry-backend/db/query/role"
	"github.com/CodeHanHan/ferry-backend/models/system"
	"github.com/CodeHanHan/ferry-backend/pkg/app"
	formRole "github.com/CodeHanHan/ferry-backend/pkg/form/role"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/pkg/sender"
	"github.com/CodeHanHan/ferry-backend/utils/stringutil"
)

var TimeNow = time.Now()

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

	newRole := system.NewRole(req.RoleName, req.Remark, system.IsAdmin(*req.IsAdmin), creator)
	if err := role.CreateRole(c, newRole); err != nil {
		if errors.Is(db.ErrDuplicateValue, err) {
			app.Error(c, app.Err_Invalid_Argument, "RoleName already exists")
			return
		}
		app.InternalServerError(c)
		return
	}

	resp := formRole.CreateRoleResponse{
		// RoleID:   newRole.RoleID,
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

// ListRole godoc
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
	var limit, offset int
	var err error
	limit, err = stringutil.String2Int(c.Request.FormValue("pageSize"))
	offset, err = stringutil.String2Int(c.Request.FormValue("pageIndex"))
	if err != nil {
		logger.ErrorParams(c, err)
		app.ErrorParams(c, err)
		return
	}

	var r system.Role
	r.RoleKey = c.Request.FormValue("roleKey")
	r.RoleName = c.Request.FormValue("roleName")
	r.Status = c.Request.FormValue("status")

	roles, count, err := r.GetPage(c, limit, offset)
	if err != nil {
		app.InternalServerError(c)
		return
	}

	app.PageOK(c, roles, int(count), offset, limit, "")
}

// GetRole godoc
// @Summary 查询角色
// @Description 根据角色id查询角色信息
// @Tags role
// @ID role-get
// @Param role_id path string true "角色id"
// @Success 200 {object} formRole.GetRoleResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Produce  json
// @Router /role/{role_id} [get]
// @Security BearerAuth
func GetRole(c *gin.Context) {
	var req formRole.GetRoleRequest
	if err := c.ShouldBindUri(&req); err != nil {
		logger.ErrorParams(c, err)
		app.ErrorParams(c, err)
		return
	}

	f := db.NewFilter().Set("role_id", req.RoleID)
	role_, err := role.GetRole(c, f)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			app.Errorf(c, app.Err_Not_found, "查询失败，未找到该记录值: %s", req.RoleID)
			return
		}
		app.InternalServerError(c)
		return
	}

	resp := formRole.GetRoleResponse{
		Role: role_,
	}

	app.OK(c, resp)
}

// UpdateRole godoc
// @Summary 更新角色
// @Description 更新角色信息
// @Tags role
// @ID role-update
// @Param role body formRole.UpdateRoleRequest true "角色id"
// @Success 200 {object} formRole.UpdateRoleResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Accept application/json
// @Produce  json
// @Router /role [put]
// @Security BearerAuth
func UpdateRole(c *gin.Context) {
	var req formRole.UpdateRoleRequest
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		logger.ErrorParams(c, err)
		app.ErrorParams(c, err)
		return
	}

	role_ := system.Role(req)

	// f := db.NewFilter().Set("role_id", role_.RoleID)
	// if _, err := role.GetRole(c, f); err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		app.Errorf(c, app.Err_Not_found, "未找到该记录：%v", role_.RoleID)
	// 		return
	// 	}
	// 	app.InternalServerError(c)
	// 	return
	// }

	f := db.NewFilter().Set("role_id", req.RoleID)
	_, err := role.GetRole(c, f)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			app.Errorf(c, app.Err_Not_found, "查询失败，未找到该记录值: %v", req.RoleID)
			return
		}
		app.InternalServerError(c)
		return
	}

	sender, _, err := sender.GetSender(c)
	if err != nil {
		logger.Error(c, err.Error())
		app.Error(c, app.Err_Unauthenticated, err)
		return
	}

	role_.UpdateBy = sender
	role_.UpdateTime = &TimeNow
	if err := role.UpdateRole(c, &role_); err != nil {
		app.InternalServerError(c)
		return
	}

	resp := formRole.UpdateRoleResponse{
		Result: "success",
	}

	app.OK(c, resp)
}
