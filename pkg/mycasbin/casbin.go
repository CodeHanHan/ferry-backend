package mycasbin

import (
	"context"

	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func SetUp(register func(*casbin.Enforcer), db *gorm.DB) error {
	adapter, err := gormadapter.NewAdapterByDBUseTableName(db, "", "casbin_rule")
	if err != nil {
		return err
	}

	e, err := casbin.NewEnforcer("deploy/config/rbac_model.conf", adapter)
	if err != nil {
		logger.Critical(context.Background(), err.Error())
		return err
	}

	if err := e.LoadPolicy(); err != nil {
		return err
	}

	register(e)

	return nil
}
