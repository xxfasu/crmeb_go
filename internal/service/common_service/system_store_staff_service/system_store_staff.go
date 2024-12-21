package system_store_staff_service

import (
	"context"
	"crmeb_go/internal/common/page"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/system_store_staff_repository"
	"crmeb_go/internal/service/common_service/system_store_service"
	"crmeb_go/internal/service/common_service/user_service"
	"crmeb_go/internal/validation"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
)

func New(
	tm repository.Transaction,
	systemStoreStaffRepo system_store_staff_repository.Repository,
	userService user_service.Service,
	systemStoreService system_store_service.Service,
) Service {
	return &service{
		tm:                   tm,
		systemStoreStaffRepo: systemStoreStaffRepo,
		userService:          userService,
		systemStoreService:   systemStoreService,
	}
}

type service struct {
	tm                   repository.Transaction
	systemStoreStaffRepo system_store_staff_repository.Repository
	userService          user_service.Service
	systemStoreService   system_store_service.Service
}

func (s *service) GetStaffList(ctx context.Context, req *validation.GetSystemStoreStaffList) (*page.CommonPageResp[*response.SystemStoreStaffResp], error) {
	resp := new(page.CommonPageResp[*response.SystemStoreStaffResp])
	systemStoreStaffRespList := make([]*response.SystemStoreStaffResp, 0)
	systemStoreStaffList, total, err := s.systemStoreStaffRepo.GetStoreStaffPageList(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(systemStoreStaffList) == 0 {
		return resp, nil
	}
	userIDList := lo.Map(systemStoreStaffList, func(item *model.SystemStoreStaff, index int) int64 {
		return item.UID
	})
	userMap := make(map[int64]*model.User)
	if len(userIDList) > 0 {
		userMap, err = s.userService.GetMapInID(ctx, userIDList)
		if err != nil {
			return nil, err
		}
	}
	storeIDList := lo.Map(systemStoreStaffList, func(item *model.SystemStoreStaff, index int) int64 {
		return item.StoreID
	})
	storeMap := make(map[int64]*model.SystemStore)
	if len(storeIDList) > 0 {
		storeMap, err = s.systemStoreService.GetMapInID(ctx, userIDList)
		if err != nil {
			return nil, err
		}
	}
	for _, storeStaff := range systemStoreStaffList {
		systemStoreStaffResp := new(response.SystemStoreStaffResp)
		err = copier.Copy(systemStoreStaffResp, storeStaff)
		if err != nil {
			return nil, err
		}
		systemStoreStaffResp.User = userMap[storeStaff.UID]
		systemStoreStaffResp.SystemStore = storeMap[storeStaff.StoreID]
		systemStoreStaffRespList = append(systemStoreStaffRespList, systemStoreStaffResp)
	}
	resp = page.RestPage(req.PageParam, systemStoreStaffRespList, total)
	return resp, nil
}
