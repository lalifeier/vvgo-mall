package service

import (
	"context"

	pb "github.com/lalifeier/vvgo-mall/api/shop/admin/v1"
)

func (s *ShopService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	return &pb.RegisterResp{}, nil
}
func (s *ShopService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	// return &pb.LoginResp{}, nil
	return s.authUsecase.Login(ctx, req)
}

func (s *ShopService) CreateSystem(ctx context.Context, req *pb.CreateSystemReq) (*pb.CreateSystemResp, error) {
	return &pb.CreateSystemResp{}, nil
}
func (s *ShopService) UpdateSystem(ctx context.Context, req *pb.UpdateSystemReq) (*pb.UpdateSystemResp, error) {
	return &pb.UpdateSystemResp{}, nil
}
func (s *ShopService) DeleteSystem(ctx context.Context, req *pb.DeleteSystemReq) (*pb.DeleteSystemResp, error) {
	return &pb.DeleteSystemResp{}, nil
}
func (s *ShopService) ListSystem(ctx context.Context, req *pb.ListSystemReq) (*pb.ListSystemResp, error) {
	return &pb.ListSystemResp{}, nil
}
func (s *ShopService) PageListSystem(ctx context.Context, req *pb.PageListSystemReq) (*pb.PageListSystemResp, error) {
	return &pb.PageListSystemResp{}, nil
}
func (s *ShopService) GetSystem(ctx context.Context, req *pb.GetSystemReq) (*pb.GetSystemResp, error) {
	return &pb.GetSystemResp{}, nil
}
func (s *ShopService) CreateRole(ctx context.Context, req *pb.CreateRoleReq) (*pb.CreateRoleResp, error) {
	return &pb.CreateRoleResp{}, nil
}
func (s *ShopService) UpdateRole(ctx context.Context, req *pb.UpdateRoleReq) (*pb.UpdateRoleResp, error) {
	return &pb.UpdateRoleResp{}, nil
}
func (s *ShopService) DeleteRole(ctx context.Context, req *pb.DeleteRoleReq) (*pb.DeleteRoleResp, error) {
	return &pb.DeleteRoleResp{}, nil
}
func (s *ShopService) ListRole(ctx context.Context, req *pb.ListRoleReq) (*pb.ListRoleResp, error) {
	return &pb.ListRoleResp{}, nil
}
func (s *ShopService) PageListRole(ctx context.Context, req *pb.PageListRoleReq) (*pb.PageListRoleResp, error) {
	return &pb.PageListRoleResp{}, nil
}
func (s *ShopService) GetRole(ctx context.Context, req *pb.GetRoleReq) (*pb.GetRoleResp, error) {
	return &pb.GetRoleResp{}, nil
}
func (s *ShopService) CreateMenu(ctx context.Context, req *pb.CreateMenuReq) (*pb.CreateMenuResp, error) {
	return &pb.CreateMenuResp{}, nil
}
func (s *ShopService) UpdateMenu(ctx context.Context, req *pb.UpdateMenuReq) (*pb.UpdateMenuResp, error) {
	return &pb.UpdateMenuResp{}, nil
}
func (s *ShopService) DeleteMenu(ctx context.Context, req *pb.DeleteMenuReq) (*pb.DeleteMenuResp, error) {
	return &pb.DeleteMenuResp{}, nil
}
func (s *ShopService) ListMenu(ctx context.Context, req *pb.ListMenuReq) (*pb.ListMenuResp, error) {
	return &pb.ListMenuResp{}, nil
}
func (s *ShopService) PageListMenu(ctx context.Context, req *pb.PageListMenuReq) (*pb.PageListMenuResp, error) {
	return &pb.PageListMenuResp{}, nil
}
func (s *ShopService) GetMenu(ctx context.Context, req *pb.GetMenuReq) (*pb.GetMenuResp, error) {
	return &pb.GetMenuResp{}, nil
}
