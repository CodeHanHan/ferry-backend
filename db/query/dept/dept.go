package dept

import (
	"context"

	"github.com/CodeHanHan/ferry-backend/db"
	modelDept "github.com/CodeHanHan/ferry-backend/models/dept"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/go-sql-driver/mysql"
)

func CreateDept(ctx context.Context, dept *modelDept.Dept) error {
	if err := db.Store.Table(modelDept.DeptTableName).Create(dept).Error; err != nil {
		logger.Error(ctx, err.Error())
		err, ok := err.(*mysql.MySQLError)
		if ok && err.Number == 1062 {
			return db.ErrDuplicateValue
		}

		return err
	}
	return nil
}

func ParentExist(ctx context.Context, parent_id string) error {
	var dept modelDept.Dept
	if err := db.Store.Table(modelDept.DeptTableName).Where("dept_id", parent_id).Take(&dept).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}
	return nil
}

func FindDeptPath(ctx context.Context, parent_id string) (string, error) {
	if parent_id == "0" {
		return "", nil
	}
	f := db.NewFilter().Set("dept_id", parent_id)
	var path string
	for {
		dept_, err := GetDept(ctx, f)
		if err != nil {
			return "", err
		}
		path = dept_.DeptName + "/" + path
		f = db.NewFilter().Set("dept_id", dept_.ParentID)
		if dept_.ParentID == "0" {
			break
		}
	}
	return path, nil
}

func DeleteDeptById(ctx context.Context, deptId string) error {
	var dept modelDept.Dept
	if err := db.Store.Table(modelDept.DeptTableName).Where(map[string]interface{}{"dept_id": deptId}).Delete(&dept).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}
	return nil
}

func SearchDept(ctx context.Context, offset, limit int) (list []*modelDept.Dept, err error) {
	if err := db.Store.Table(modelDept.DeptTableName).Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}

	return
}

func GetDept(ctx context.Context, f *db.Filter) (dept *modelDept.Dept, err error) {
	if err := db.Store.Table(modelDept.DeptTableName).Where(f.Params).Take(&dept).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}

	return
}

func UpdateDept(ctx context.Context, dept *modelDept.Dept) error {
	if err := db.Store.Table(modelDept.DeptTableName).Updates(dept).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	return nil

}

// func UpdateDeptPath(ctx context.Context)
