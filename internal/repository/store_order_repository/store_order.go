package store_order_repository

import (
	"context"
	"crmeb_go/constants"
	"crmeb_go/internal/common/data"
	"crmeb_go/internal/repository/gen"
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func New(
	db *gorm.DB,
) Repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *gorm.DB
}

func (r *repository) FindOrderNumByDate(ctx context.Context, start int64, end int64) (data int64, err error) {
	storeOrder := gen.Q.StoreOrder

	return storeOrder.WithContext(ctx).Select(storeOrder.ID).
		Where(
			storeOrder.CreatedAt.Between(start, end),
			storeOrder.Paid.Eq(constants.StoreOrderPaidValid),
		).Count()
}

func (r *repository) FindPayOrderAmountByDate(ctx context.Context, start int64, end int64) (payPrice decimal.Decimal, err error) {
	storeOrder := gen.Q.StoreOrder
	first, err := storeOrder.WithContext(ctx).Select(storeOrder.PayPrice).
		Where(
			storeOrder.CreatedAt.Between(start, end),
			storeOrder.Paid.Eq(constants.StoreOrderPaidValid),
		).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return payPrice, err
	}
	if first != nil {
		payPrice = first.PayPrice
	}
	return payPrice, nil
}

func (r *repository) FindOrderGroupByDate(ctx context.Context, start int64, end int64) ([]*data.StoreOrderEveryDate, error) {
	storeOrder := gen.Q.StoreOrder
	return storeOrder.WithContext(ctx).QueryOrderGroupByDate(&data.DateCondition{Start: start, End: end})
}
