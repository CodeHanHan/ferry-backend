package mycasbin

import (
	"github.com/CodeHanHan/ferry-backend/pkg/pi"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func Casbin() (*casbin.Enforcer, error) {
	adapter, err := gormadapter.NewAdapterByDBUseTableName(pi.Global.Mysql, "", "casbin_rule")
	if err != nil {
		return nil, err
	}

	e, err := casbin.NewEnforcer("deploy/config/rbac_model.conf", adapter)
	if err != nil {
		return nil, err
	}

	if err := e.LoadPolicy(); err != nil {
		return nil, err
	}

	return e, nil
}
