package data

import (
	"crmeb_go/internal/model"
	"github.com/shopspring/decimal"
)

type DateParams struct {
	Start          int64 // 今天0点间戳
	End            int64 // 今天结束时间戳
	YesterdayStart int64 // 咋天0点时间戳
	YesterdayEnd   int64 // 咋天24点时间戳
}

type DateResult struct {
	NowData       int64 // 今天结果
	YesterdayData int64 // 咋天结果
	NowSale       decimal.Decimal
	YesterdaySale decimal.Decimal
}

type DateCondition struct {
	Start int64
	End   int64
}

type StoreOrderEveryDate struct {
	model.StoreOrder
	EveryDate string `gorm:"column:every_date;type:varchar(32);not null;comment:每天日期" json:"every_date"` // 每天日期
}

type UserEveryDate struct {
	model.User
	EveryDate string `gorm:"column:every_date;type:varchar(32);not null;comment:每天日期" json:"every_date"` // 每天日期
}
