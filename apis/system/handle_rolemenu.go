package system

import (
	"github.com/CodeHanHan/ferry-backend/models/system"
	"github.com/CodeHanHan/ferry-backend/pkg/app"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

// @Summary RoleMenu列表数据
// @Description 获取JSON
// @Tags 角色菜单
// @Param RoleId query string false "RoleId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/rolemenu [get]
// @Security Bearer
func GetRoleMenu(c *gin.Context) {
	var rm system.BindRoleMenu
	if err := c.ShouldBind(&rm); err != nil {
		logger.ErrorParams(c, err)
		app.ErrorParams(c, err)
		return
	}

	rms, err := rm.Get(c)
	if err != nil {
		app.InternalServerError(c)
		return
	}

	var resp app.Response
	resp.Data = rms
	app.OK(c, resp.ReturnOK())
}

// @Summary 删除用户菜单数据
// @Description 删除数据
// @Tags 角色菜单
// @Param id path string true "id"
// @Param menu_id query string false "menu_id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/rolemenu/{id} [delete]
func DeleteRoleMenu(c *gin.Context) {
	var rm system.BindRoleMenu

	id := c.Param("id")
	menuId := c.Request.FormValue("menu_id")
	if err := rm.Delete(c, id, menuId); err != nil {
		app.InternalServerError(c)
		return
	}

	var resp app.Response
	resp.Msg = "success"
	app.OK(c, resp)
}
