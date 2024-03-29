package biz

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/lalifeier/vvgo-mall/pkg/util/convert"
)

type Role struct {
	Id int64
}

type RoleListReq struct {
}

type RolePageListReq struct {
	PageNum  int64
	PageSize int64
}

type RolePageListResp struct {
	TotalPages  int64
	CurrentPage int64
	PageSize    int64
	Data        []*Role
}

type RoleRepo interface {
	CreateRole(ctx context.Context, role *Role) (int64, error)
	UpdateRole(ctx context.Context, role *Role) error
	DeleteRole(ctx context.Context, id int64) error
	GetRole(ctx context.Context, id int64) (*Role, error)
	ListRole(ctx context.Context, req *RoleListReq) ([]*Role, error)
	PageListRole(ctx context.Context, req *RolePageListReq) (*RolePageListResp, error)
}

type RoleUseCase struct {
	repo     RoleRepo
	log      *log.Helper
	enforcer *casbin.Enforcer
}

func NewRoleUseCase(repo RoleRepo, enforcer *casbin.Enforcer, logger log.Logger) *RoleUseCase {
	return &RoleUseCase{repo: repo, enforcer: enforcer, log: log.NewHelper(logger)}
}

func (uc *RoleUseCase) CreateRole(ctx context.Context, role *Role) (int64, error) {
	id, err := uc.repo.CreateRole(ctx, role)
	if err != nil {
		return 0, err
	}

	// uc.enforcer.AddPermissionForUser()

	return id, nil
}

func (uc *RoleUseCase) UpdateRole(ctx context.Context, role *Role) error {
	err := uc.repo.UpdateRole(ctx, role)
	if err != nil {
		return err
	}

	// uc.enforcer.DeleteRole(convert.Int64ToString(role.Id))
	// uc.enforcer.AddPermissionForUser(convert.Int64ToString(role.Id))

	return nil
}

func (uc *RoleUseCase) DeleteRole(ctx context.Context, id int64) error {
	err := uc.repo.DeleteRole(ctx, id)
	if err != nil {
		return err
	}

	uc.enforcer.DeleteRole(convert.Int64ToString(id))
	return nil
}

func (uc *RoleUseCase) GetRole(ctx context.Context, id int64) (*Role, error) {
	return uc.repo.GetRole(ctx, id)
}

func (uc *RoleUseCase) ListRole(ctx context.Context, req *RoleListReq) ([]*Role, error) {
	return uc.repo.ListRole(ctx, req)
}

func (uc *RoleUseCase) PageListRole(ctx context.Context, req *RolePageListReq) (*RolePageListResp, error) {
	return uc.repo.PageListRole(ctx, req)
}

// func (uc *RoleUseCase) compareRoleRoles(ctx context.Context, oldRoleRoles, newRoleRoles RoleRoles) (addList, delList RoleRoles) {
// 	return
// }
