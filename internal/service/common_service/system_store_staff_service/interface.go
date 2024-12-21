package system_store_staff_service

import (
	"context"
	"crmeb_go/internal/common/page"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/validation"
)

type Service interface {
	GetStaffList(ctx context.Context, req *validation.GetSystemStoreStaffList) (*page.CommonPageResp[*response.SystemStoreStaffResp], error)
}
