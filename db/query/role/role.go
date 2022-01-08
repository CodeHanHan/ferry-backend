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
