package role

import (
	"context"

	"github.com/CodeHanHan/ferry-backend/db"
	modelRole "github.com/CodeHanHan/ferry-backend/models/role"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/go-sql-driver/mysql"
)

func CreateRole(ctx context.Context, role *modelRole.Role) error {
	if err := db.Store.Table(modelRole.RoleTableName).Create(role).Error; err != nil {
		logger.Error(ctx, err.Error())
		err, ok := err.(*mysql.MySQLError)
		if ok && err.Number == 1062 {
			return db.ErrDuplicateValue
		}

		return err
	}

	return nil
}

func DeleteRoleById(ctx context.Context, roleId string) error {
	var role modelRole.Role
	if err := db.Store.Table(modelRole.RoleTableName).Where(map[string]interface{}{"role_id": roleId}).Delete(&role).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	return nil
}

func SearchRole(ctx context.Context, offset, limit int) (list []*modelRole.Role, err error) {
	if err := db.Store.Table(modelRole.RoleTableName).Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}

	return
}

func GetRole(ctx context.Context, f *db.Filter) (role *modelRole.Role, err error) {
	if err := db.Store.Table(modelRole.RoleTableName).Where(f.Params).Take(&role).Error; err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}

	return
}

func UpdateRole(ctx context.Context, role *modelRole.Role) error {
	if err := db.Store.Table(modelRole.RoleTableName).Updates(role).Error; err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	return nil
}
