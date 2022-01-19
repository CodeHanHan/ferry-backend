package menu

import (
	modelMenu "github.com/CodeHanHan/ferry-backend/models/menu"
	"github.com/CodeHanHan/ferry-backend/pkg/app"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/pkg/sender"
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
	var menu modelMenu.Menu
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
