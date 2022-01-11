package dept

import (
	"time"

	"github.com/CodeHanHan/ferry-backend/utils/idutil"
)

const DeptTableName = "dept"

type DeptStatus int

const (
	DeptDeactivated DeptStatus = 0
	DeptActive      DeptStatus = 1
)

type Dept struct {
	DeptID     string     `gorm:"column:dept_id;primary_key" json:"dept_id"`
	ParentID   string     `gorm:"column:parent_id" json:"parent_id"`
	DeptPath   string     `gorm:"column:dept_path" json:"dept_path"`
	DeptName   string     `gorm:"column:dept_name" json:"dept_name"`
	DeptSort   int        `gorm:"column:dept_sort" json:"dept_sort"`
	Leader     string     `gorm:"column:leader" json:"leader"`
	Phone      string     `gorm:"column:phone" json:"phone"`
	Email      string     `gorm:"column:email" json:"email"`
	Status     int        `gorm:"column:status" json:"status"`
	CreateBy   string     `gorm:"column:create_by" json:"create_by"`
	UpdateBy   string     `gorm:"column:update_by" json:"update_by"`
	CreateTime *time.Time `gorm:"column:create_time" json:"create_time" default:"2000-01-01 00:00:00"`
	UpdateTime *time.Time `gorm:"column:update_time" json:"update_time" default:"2000-01-01 00:00:00"`
	DeleteTime *time.Time `gorm:"column:delete_time" json:"delete_time" default:"2000-01-01 00:00:00"`
}

func NewDept(deptname string, parent_id string, create_by string) *Dept {
	return &Dept{
		DeptID:   idutil.GetId("dept"),
		DeptName: deptname,
		ParentID: parent_id,
		Status:   int(DeptActive),
		CreateBy: create_by,
	}

}
