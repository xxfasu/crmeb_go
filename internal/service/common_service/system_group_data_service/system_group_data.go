package system_group_data_service

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/system_group_data_repository"
	"crmeb_go/internal/validation"
	"crmeb_go/pkg/util"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
)

func New(
	redisClient redis.UniversalClient,
	tm repository.Transaction,
	systemGroupDataRepo system_group_data_repository.Repository,
) Service {
	return &service{
		redisClient:         redisClient,
		tm:                  tm,
		systemGroupDataRepo: systemGroupDataRepo,
	}
}

type service struct {
	redisClient         redis.UniversalClient
	tm                  repository.Transaction
	systemGroupDataRepo system_group_data_repository.Repository
}

func (s *service) GetList(ctx context.Context, req validation.SystemGroupDataSearchReq) ([]*model.SystemGroupData, error) {
	list, _, err := s.systemGroupDataRepo.GetGroupDataPageList(ctx, req)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *service) GetListByGID(ctx context.Context, gid int64) ([]any, error) {
	var systemGroupDataSearchReq validation.SystemGroupDataSearchReq
	systemGroupDataSearchReq.GID = gid
	systemGroupDataSearchReq.Status = 1
	systemGroupDataSearchReq.PageParam = util.DefaultPageParams()
	list, err := s.GetList(ctx, systemGroupDataSearchReq)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("暂无数据")
	}
	result := make([]any, 0, len(list))
	for _, systemGroupData := range list {
		mapData := make(map[string]any)
		json.Unmarshal([]byte(systemGroupData.Value), &mapData)
		var systemFormItemCheckReqList []*validation.SystemFormItemCheckReq
		fields := mapData["fields"]
		marshal, _ := json.Marshal(fields)
		json.Unmarshal(marshal, &systemFormItemCheckReqList)
		if len(systemFormItemCheckReqList) == 0 {
			continue
		}
		data := make(map[string]any)
		for _, systemFormItemCheckReq := range systemFormItemCheckReqList {
			if systemFormItemCheckReq.Name != "pic" {
				data[systemFormItemCheckReq.Name] = systemFormItemCheckReq.Value
			} else {
				data[systemFormItemCheckReq.Name] = "http://localhost:7788/" + systemFormItemCheckReq.Value
			}
		}
		data["id"] = systemGroupData.ID
		result = append(result, data)
	}
	return result, nil
}
