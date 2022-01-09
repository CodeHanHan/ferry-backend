package dept

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"

	"github.com/CodeHanHan/ferry-backend/db"
	dept "github.com/CodeHanHan/ferry-backend/db/query/dept"
	modelDept "github.com/CodeHanHan/ferry-backend/models/dept"
	"github.com/CodeHanHan/ferry-backend/pkg/app"
	formDept "github.com/CodeHanHan/ferry-backend/pkg/form/dept"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/pkg/sender"
)

// CreateDept godoc
// @Summary 创建部门
// @Description 根据一个DeptName创建部门，DeptName 要求不能重复
// @Tags dept
// @ID dept-create
// @Param dept body formDept.CreateDeptRequest true "参数不可空"
// @Success 200 {object} formDept.CreateDeptResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Accept application/json
// @Produce  json
// @Router /dept [post]
// @Security BearerAuth
func CreateDept(c *gin.Context) {
	var req formDept.CreateDeptRequest
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

	newDept := modelDept.NewDept(req.DeptName, req.ParentID, creator)
	if err := dept.CreateDept(c, newDept); err != nil {
		if errors.Is(db.ErrDuplicateValue, err) {
			app.Error(c, app.Err_Invalid_Argument, "DeptName already exists")
			return
		}
		app.InternalServerError(c)
		return
	}

	resp := formDept.CreateDeptResponse{
		DeptID:   newDept.DeptID,
		DeptName: newDept.DeptName,
	}

	app.OK(c, resp)
}

// DeleteDept godoc
// @Summary 删除部门
// @Description 删除部门，幂等操作
// @Tags dept
// @ID dept-delete
// @Param dept_id path string true "部门唯一id"
// @Success 200 {object} formDept.DeleteDeptResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Accept application/json
// @Produce  json
// @Router /dept/{dept_id} [delete]
// @Security BearerAuth
func DeleteDept(c *gin.Context) {
	var req formDept.DeleteDeptRequest
	if err := c.ShouldBindUri(&req); err != nil {
		logger.ErrorParams(c, err)
		app.ErrorParams(c, err)
		return
	}

	if err := dept.DeleteDeptById(c, req.DeptID); err != nil {
		app.InternalServerError(c)
		return
	}

	app.OK(c, formDept.DeleteDeptResponse{
		Result: "success",
	})

}

//ListDept godoc
// @Summary 查询部门列表
// @Description 根据offset和limit查询部门列表
// @Tags dept
// @ID dept-list
// @Param offset query int true "偏移"
// @Param limit query int true "限制"
// @Success 200 {object} formDept.ListDeptResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Produce  json
// @Router /dept [get]
// @Security BearerAuth
func ListDpet(c *gin.Context) {
	var req formDept.ListDeptRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.ErrorParams(c, err)
		app.ErrorParams(c, err)
		return
	}

	list, err := dept.SearchDept(c, *req.Offset, req.Limit)
	if err != nil {
		app.InternalServerError(c)
		return
	}

	resp := formDept.ListDeptResponse{
		Dept:   list,
		Length: len(list),
	}

	app.OK(c, resp)
}

//GetDept godoc
// @Summary 查询部门
// @Description 根据部门id查询部门信息
// @Tags dept
// @ID dept-get
// @Param dept_id path string true "部门id"
// @Success 200 {object} formDept.GetDeptResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Produce  json
// @Router /dept/{dept_id} [get]
// @Security BearerAuth
func GetDept(c *gin.Context) {
	var req formDept.GetDeptRequest
	if err := c.ShouldBindUri(&req); err != nil {
		logger.ErrorParams(c, err)
		app.ErrorParams(c, err)
		return
	}

	f := db.NewFilter().Set("dept_id", req.DeptID)
	dept, err := dept.GetDept(c, f)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			app.Errorf(c, app.Err_Not_found, "查询失败，未找到该记录值: %s", req.DeptID)
			return
		}
		app.InternalServerError(c)
		return
	}

	resp := formDept.GetDeptResponse{
		Dept: dept,
	}

	app.OK(c, resp)
}

//UpdateDept godoc
// @Summary 更新部门
// @Description 更新部门信息
// @Tags dept
// @ID dept-update
// @Param dept body formDept.UpdateDeptRequest true "部门id"
// @Success 200 {object} formDept.UpdateDeptResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Accept application/json
// @Produce  json
// @Router /dept [put]
// @Security BearerAuth
func UpdateDept(c *gin.Context) {
	var req formDept.UpdateDeptRequest
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		logger.ErrorParams(c, err)
		app.ErrorParams(c, err)
		return
	}

	dept_ := modelDept.Dept(req)

	f := db.NewFilter().Set("dept_id", req.DeptID)
	_, err := dept.GetDept(c, f)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			app.Errorf(c, app.Err_Not_found, "查询失败，未找到该记录值: %s", req.DeptID)
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

	var TimeNow = time.Now()
	dept_.UpdateBy = sender
	dept_.UpdateTime = &TimeNow
	if err := dept.UpdateDept(c, &dept_); err != nil {
		app.InternalServerError(c)
		return
	}

	resp := formDept.UpdateDeptResponse{
		Result: "success",
	}

	app.OK(c, resp)
}
