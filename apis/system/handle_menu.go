package system

import (
	"github.com/CodeHanHan/ferry-backend/models/system"
	"github.com/CodeHanHan/ferry-backend/pkg/app"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/pkg/sender"
	"github.com/CodeHanHan/ferry-backend/utils/stringutil"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary 创建菜单
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/x-www-form-urlencoded
// @Product application/x-www-form-urlencoded
// @Param menuName formData string true "menuName"
// @Param Path formData string false "Path"
// @Param Action formData string true "Action"
// @Param Permission formData string true "Permission"
// @Param ParentId formData string true "ParentId"
// @Param IsDel formData string true "IsDel"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /menu [post]
// @Security Bearer
func CreateMenu(c *gin.Context) {
	var menu system.Menu
	if err := c.BindWith(&menu, binding.JSON); err != nil {
		logger.ErrorParams(c, err)
		app.ErrorParams(c, err)
		return
	}

	sender, _, err := sender.GetSender(c)
	if err != nil {
		app.Error(c, app.Err_Unauthenticated, "用户身份获取失败")
		return
	}

	menu.CreateBy = sender

	res, err := menu.Create(c)
	if err != nil {
		app.InternalServerError(c)
		return
	}

	app.AdaptOK(c, res, "")
}

// @Summary 修改菜单
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/x-www-form-urlencoded
// @Product application/x-www-form-urlencoded
// @Param id path int true "id"
// @Param data body system.Menu true "body"
// @Success 200 {string} string	"{"code": 200, "message": "修改成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "修改失败"}"
// @Router /api/v1/menu/{id} [put]
// @Security Bearer
func UpdateMenu(c *gin.Context) {
	var menu system.Menu
	if err := c.BindWith(&menu, binding.JSON); err != nil {
		logger.ErrorParams(c, err)
		app.ErrorParams(c, err)
		return
	}

	sender, _, err := sender.GetSender(c)
	if err != nil {
		logger.Error(c, err.Error())
		app.Error(c, app.Err_Unauthenticated, err)
		return
	}

	menu.UpdateBy = sender

	_, err = menu.Update(c, menu.MenuID)
	if err != nil {
		app.InternalServerError(c)
		return
	}

	app.AdaptOK(c, "", "success")
}

// @Summary 删除菜单
// @Description 删除数据
// @Tags 菜单
// @Param id path int true "id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/menu/{id} [delete]
func DeleteMenu(c *gin.Context) {
	var menu system.Menu
	id, err := stringutil.String2Int(c.Param("id"))
	if err != nil {
		logger.Error(c, err.Error())
		app.Error(c, app.Err_Invalid_Argument, "")
		return
	}

	if err := menu.DeleteMenu(c, id); err != nil {
		app.InternalServerError(c)
		return
	}

	app.AdaptOK(c, "", "success")
}

// @Summary 获取菜单
// @Description 根据id获取菜单信息
// @Tags 菜单
// @Param menuName query string false "menuName"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/menu [get]
// @Security Bearer
func GetMenu(c *gin.Context) {
	var menu system.Menu
	id, err := stringutil.String2Int(c.Param("id"))
	if err != nil {
		logger.Error(c, err.Error())
		app.Error(c, app.Err_Invalid_Argument, "")
		return
	}

	menu_, err := menu.GetMenuById(c, id)
	if err != nil {
		app.InternalServerError(c) // TODO not found
		return
	}

	app.AdaptOK(c, menu_, "")
}

// @Summary Menu列表数据
// @Description 获取JSON
// @Tags 菜单
// @Param menuName query string false "menuName"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/menulist [get]
// @Security Bearer
func GetMenuList(c *gin.Context) {
	var menu system.Menu
	menu.MenuName = c.Request.FormValue("menuName")
	menu.Visible = c.Request.FormValue("visible")
	menu.Title = c.Request.FormValue("title")

	var res []*system.Menu
	var err error
	if menu.Title == "" {
		res, err = menu.SetMenu(c)
	} else {
		res, err = menu.GetPage(c)
	}
	if err != nil {
		app.InternalServerError(c)
		return
	}

	app.AdaptOK(c, res, "")
}

// func GetMenuTreeRoleSelect(c *gin.Context) {

// }
