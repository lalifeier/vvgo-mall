package biz

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/lalifeier/vvgo-mall/pkg/util/convert"
)

type UserRole struct {
	Id     int64
	UserId int64
	RoleId int64
}

type UserRoles []*UserRole

type User struct {
	Id        int64
	UserRoles UserRoles
}

type UserListReq struct {
}

type UserPageListReq struct {
	PageNum  int64
	PageSize int64
}

type UserPageListResp struct {
	TotalPages  int64
	CurrentPage int64
	PageSize    int64
	Data        []*User
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (int64, error)
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id int64) error
	GetUser(ctx context.Context, id int64) (*User, error)
	ListUser(ctx context.Context, req *UserListReq) ([]*User, error)
	PageListUser(ctx context.Context, req *UserPageListReq) (*UserPageListResp, error)
}

type UserUsecase struct {
	repo     UserRepo
	log      *log.Helper
	enforcer *casbin.Enforcer
}

func NewUserUsecase(repo UserRepo, enforcer *casbin.Enforcer, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, enforcer: enforcer, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, user *User) (int64, error) {
	id, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}
	for _, item := range user.UserRoles {
		uc.enforcer.AddRoleForUser(convert.Int64ToString(item.UserId), convert.Int64ToString(item.RoleId))
	}
	return id, nil
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, user *User) error {
	err := uc.repo.UpdateUser(ctx, user)
	if err != nil {
		return err
	}

	addUserRoles := make([]*UserRole, 0)
	delUserRoles := make([]*UserRole, 0)
	for _, item := range addUserRoles {
		uc.enforcer.AddRoleForUser(convert.Int64ToString(item.UserId), convert.Int64ToString(item.RoleId))
	}

	for _, item := range delUserRoles {
		uc.enforcer.DeleteRoleForUser(convert.Int64ToString(item.UserId), convert.Int64ToString(item.RoleId))
	}

	return nil
}

func (uc *UserUsecase) DeleteUser(ctx context.Context, id int64) error {
	err := uc.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	uc.enforcer.DeleteUser(convert.Int64ToString(id))
	return nil
}

func (uc *UserUsecase) GetUser(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUsecase) ListUser(ctx context.Context, req *UserListReq) ([]*User, error) {
	return uc.repo.ListUser(ctx, req)
}

func (uc *UserUsecase) PageListUser(ctx context.Context, req *UserPageListReq) (*UserPageListResp, error) {
	return uc.repo.PageListUser(ctx, req)
}

func (uc *UserUsecase) compareUserRoles(ctx context.Context, oldUserRoles, newUserRoles UserRoles) (addList, delList UserRoles) {
	return
}
