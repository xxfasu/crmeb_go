package store_order_repository

import (
	"context"
	"crmeb_go/internal/common/data"
	"github.com/shopspring/decimal"
)

type Reader interface {
	FindOrderNumByDate(ctx context.Context, start int64, end int64) (int64, error)
	FindPayOrderAmountByDate(ctx context.Context, start int64, end int64) (payPrice decimal.Decimal, err error)
	FindOrderGroupByDate(ctx context.Context, start int64, end int64) ([]*data.StoreOrderEveryDate, error)
}

type Writer interface {
}

type Repository interface {
	Reader
	Writer
}
