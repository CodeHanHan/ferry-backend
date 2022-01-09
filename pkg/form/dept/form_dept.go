package dept

import "github.com/CodeHanHan/ferry-backend/models/dept"

type CreateDeptRequest struct {
	DeptName string `json:"dept_name" form:"dept_name" binding:"required"`
	ParentID string `json:"parent_id" form:"parent_id" binding:"required"`
}

type CreateDeptResponse struct {
	DeptID   string `json:"dept_id"`
	DeptName string `json:"dept_name"`
}

type DeleteDeptRequest struct {
	DeptID string `json:"dept_id" uri:"dept_id" binding:"required"`
}

type DeleteDeptResponse struct {
	Result string `json:"result"`
}

type ListDeptRequest struct {
	Offset *int `json:"offset" form:"offset" binding:"required"`
	Limit  int  `json:"limit" form:"limit" binding:"required"`
}

type ListDeptResponse struct {
	Dept  []*dept.Dept
	Length int
}

type GetDeptRequest struct {
	DeptID string `json:"dept_id" uri:"dept_id" binding:"required"`
}

type GetDeptResponse struct {
	Dept *dept.Dept
}

type UpdateDeptRequest dept.Dept

type UpdateDeptResponse struct {
	Result string `json:"result"`
}
