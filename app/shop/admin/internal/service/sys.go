package service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/biz"
	pb "github.com/lalifeier/vvgo-mall/gen/api/shop/admin/v1"
)

func (s *ShopService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserResp, error) {
	return &pb.CreateUserResp{}, nil
}
func (s *ShopService) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	return &pb.UpdateUserResp{}, nil
}
func (s *ShopService) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserResp, error) {
	return &pb.DeleteUserResp{}, nil
}
func (s *ShopService) ListUser(ctx context.Context, req *pb.ListUserReq) (*pb.ListUserResp, error) {
	return &pb.ListUserResp{}, nil
}
func (s *ShopService) PageListUser(ctx context.Context, req *pb.PageListUserReq) (*pb.PageListUserResp, error) {
	return &pb.PageListUserResp{}, nil
}
func (s *ShopService) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserResp, error) {
	return &pb.GetUserResp{}, nil
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

func (s *ShopService) CreateDictData(ctx context.Context, req *pb.CreateDictDataReq) (*pb.CreateDictDataResp, error) {
	id, err := s.dictDataUseCase.CreateDictData(ctx, &biz.DictData{
		DictTypeId: req.DictTypeId,
		Label:      req.Label,
		Value:      req.Value,
		Status:     int8(req.Status),
		Remark:     req.Remark,
		Sort:       int8(req.Sort),
		IsDefault:  int8(req.IsDefault),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateDictDataResp{
		Id: id,
	}, nil
}
func (s *ShopService) UpdateDictData(ctx context.Context, req *pb.UpdateDictDataReq) (*pb.UpdateDictDataResp, error) {
	err := s.dictDataUseCase.UpdateDictData(ctx, &biz.DictData{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateDictDataResp{}, nil
}
func (s *ShopService) DeleteDictData(ctx context.Context, req *pb.DeleteDictDataReq) (*pb.DeleteDictDataResp, error) {
	err := s.dictDataUseCase.DeleteDictData(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteDictDataResp{}, nil
}
func (s *ShopService) ListDictData(ctx context.Context, req *pb.ListDictDataReq) (*pb.ListDictDataResp, error) {
	rv, err := s.dictDataUseCase.ListDictData(ctx, &biz.DictDataListReq{})
	if err != nil {
		return nil, err
	}

	dictdatas := make([]*pb.DictData, 0)
	err = copier.Copy(&dictdatas, &rv)
	if err != nil {
		return nil, err
	}

	return &pb.ListDictDataResp{
		Data: dictdatas,
	}, nil
}
func (s *ShopService) PageListDictData(ctx context.Context, req *pb.PageListDictDataReq) (*pb.PageListDictDataResp, error) {
	pos, err := s.dictDataUseCase.PageListDictData(ctx, &biz.DictDataPageListReq{PageNum: req.PageNum, PageSize: req.PageSize})
	if err != nil {
		return nil, err
	}

	dictdatas := make([]*pb.DictData, 0)
	err = copier.Copy(&dictdatas, &pos.Data)
	if err != nil {
		return nil, err
	}

	return &pb.PageListDictDataResp{
		TotalPages:  pos.TotalPages,
		CurrentPage: req.PageSize,
		PageSize:    req.PageNum,
		Data:        dictdatas,
	}, nil
}
func (s *ShopService) GetDictData(ctx context.Context, req *pb.GetDictDataReq) (*pb.GetDictDataResp, error) {
	au, err := s.dictDataUseCase.GetDictData(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	dictdata := pb.GetDictDataResp{}
	err = copier.Copy(&dictdata, au)
	if err != nil {
		return nil, err
	}

	return &dictdata, nil
}
