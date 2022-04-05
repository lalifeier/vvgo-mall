package data

import (
	"context"

	"fmt"

	casbinModel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/go-kratos/kratos/v2/log"
)

var _ persist.Adapter = (*casbinAdapter)(nil)

type casbinAdapter struct {
	data *Data
	log  *log.Helper
}

func NewCasbinAdapter(data *Data, logger log.Logger) persist.Adapter {
	return &casbinAdapter{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "shop-admin/data")),
	}
}

// p,role_id,path,method
func (a *casbinAdapter) loadRoleApiPolicy(ctx context.Context, m casbinModel.Model) error {
	// pos, err := a.data.db.Role.Query().All(ctx)
	// if err != nil {
	// 	return err
	// }
	// if len(pos) == 0 {
	// 	return nil
	// }
	// for _, po := range pos {
	// 	line := fmt.Sprintf("p,%d,%s,%s", item.ID, permission.Path, permission.Action)
	// 	persist.LoadPolicyLine(line, m)
	// }
	return nil
}

// g,user_id,role_id
func (a *casbinAdapter) loadUserRolePolicy(ctx context.Context, m casbinModel.Model) error {
	// users, err := a.data.db.User.Query().All(ctx)
	// if err != nil {
	// 	return err
	// }
	// if len(users) == 0 {
	// 	return nil
	// }
	// for _, user := range users {
	// 	fmt.Println(user.ID)
	// 	roles, err := user.QueryRoles().All(ctx)
	// 	if err == nil {
	// 		for _, role := range roles {
	// 			line := fmt.Sprintf("g,%d,%d", user.ID, role.ID)
	// 			persist.LoadPolicyLine(line, m)
	// 		}
	// 	}
	// }
	userRoles, err := a.data.db.UserRole.Query().All(ctx)
	if err != nil {
		return err
	}
	for _, userRole := range userRoles {
		line := fmt.Sprintf("g,%d,%d", userRole.UserID, userRole.RoleID)
		println(line)
		persist.LoadPolicyLine(line, m)
	}
	return nil
}

// LoadPolicy loads all policy rules from the storage.
func (a *casbinAdapter) LoadPolicy(model casbinModel.Model) error {
	ctx := context.Background()
	// err := a.loadRolePolicy(ctx, model)
	// if err != nil {
	// 	return err
	// }
	err := a.loadUserRolePolicy(ctx, model)
	if err != nil {
		return err
	}
	return nil
}

// SavePolicy saves all policy rules to the storage.
func (a *casbinAdapter) SavePolicy(model casbinModel.Model) error {
	return nil
}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (a *casbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (a *casbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (a *casbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}
