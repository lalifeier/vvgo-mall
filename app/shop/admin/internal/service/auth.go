package service

import (
	"context"

	pb "github.com/lalifeier/vvgo-mall/api/shop/admin/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) CreateSystem(ctx context.Context, req *pb.CreateSystemReq) (*pb.CreateSystemResp, error) {
	return &pb.CreateSystemResp{}, nil
}
func (s *AuthService) UpdateSystem(ctx context.Context, req *pb.UpdateSystemReq) (*pb.UpdateSystemResp, error) {
	return &pb.UpdateSystemResp{}, nil
}
func (s *AuthService) DeleteSystem(ctx context.Context, req *pb.DeleteSystemReq) (*pb.DeleteSystemResp, error) {
	return &pb.DeleteSystemResp{}, nil
}
func (s *AuthService) ListSystem(ctx context.Context, req *pb.ListSystemReq) (*pb.ListSystemResp, error) {
	return &pb.ListSystemResp{}, nil
}
func (s *AuthService) PageListSystem(ctx context.Context, req *pb.PageListSystemReq) (*pb.PageListSystemResp, error) {
	return &pb.PageListSystemResp{}, nil
}
func (s *AuthService) GetSystem(ctx context.Context, req *pb.GetSystemReq) (*pb.GetSystemResp, error) {
	return &pb.GetSystemResp{}, nil
}
func (s *AuthService) CreateRole(ctx context.Context, req *pb.CreateRoleReq) (*pb.CreateRoleResp, error) {
	return &pb.CreateRoleResp{}, nil
}
func (s *AuthService) UpdateRole(ctx context.Context, req *pb.UpdateRoleReq) (*pb.UpdateRoleResp, error) {
	return &pb.UpdateRoleResp{}, nil
}
func (s *AuthService) DeleteRole(ctx context.Context, req *pb.DeleteRoleReq) (*pb.DeleteRoleResp, error) {
	return &pb.DeleteRoleResp{}, nil
}
func (s *AuthService) ListRole(ctx context.Context, req *pb.ListRoleReq) (*pb.ListRoleResp, error) {
	return &pb.ListRoleResp{}, nil
}
func (s *AuthService) PageListRole(ctx context.Context, req *pb.PageListRoleReq) (*pb.PageListRoleResp, error) {
	return &pb.PageListRoleResp{}, nil
}
func (s *AuthService) GetRole(ctx context.Context, req *pb.GetRoleReq) (*pb.GetRoleResp, error) {
	return &pb.GetRoleResp{}, nil
}
func (s *AuthService) CreateMenu(ctx context.Context, req *pb.CreateMenuReq) (*pb.CreateMenuResp, error) {
	return &pb.CreateMenuResp{}, nil
}
func (s *AuthService) UpdateMenu(ctx context.Context, req *pb.UpdateMenuReq) (*pb.UpdateMenuResp, error) {
	return &pb.UpdateMenuResp{}, nil
}
func (s *AuthService) DeleteMenu(ctx context.Context, req *pb.DeleteMenuReq) (*pb.DeleteMenuResp, error) {
	return &pb.DeleteMenuResp{}, nil
}
func (s *AuthService) ListMenu(ctx context.Context, req *pb.ListMenuReq) (*pb.ListMenuResp, error) {
	return &pb.ListMenuResp{}, nil
}
func (s *AuthService) PageListMenu(ctx context.Context, req *pb.PageListMenuReq) (*pb.PageListMenuResp, error) {
	return &pb.PageListMenuResp{}, nil
}
func (s *AuthService) GetMenu(ctx context.Context, req *pb.GetMenuReq) (*pb.GetMenuResp, error) {
	return &pb.GetMenuResp{}, nil
}
