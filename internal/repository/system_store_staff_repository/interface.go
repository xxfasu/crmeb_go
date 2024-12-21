package system_store_staff_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/validation"
)

type Reader interface {
	GetStoreStaffPageList(ctx context.Context, condition *validation.GetSystemStoreStaffList) ([]*model.SystemStoreStaff, int64, error)
}

type Writer interface {
}

type Repository interface {
	Reader
	Writer
}
