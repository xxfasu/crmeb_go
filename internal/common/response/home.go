package response

import (
	"github.com/iancoleman/orderedmap"
	"github.com/shopspring/decimal"
)

type HomeRate struct {
	Sale                decimal.Decimal `json:"sales"`               // 今日销售额
	YesterdaySale       decimal.Decimal `json:"yesterdaySales"`      // 昨日销售额
	PageViews           int64           `json:"pageviews"`           // 今日访问量
	YesterdayPageViews  int64           `json:"yesterdayPageviews"`  // 昨日访问量
	OrderNum            int64           `json:"orderNum"`            // 今日订单量
	YesterdayOrderNum   int64           `json:"yesterdayOrderNum"`   // 昨日订单量
	NewUserNum          int64           `json:"newUserNum"`          // 今日新增用户
	YesterdayNewUserNum int64           `json:"yesterdayNewUserNum"` // 昨日新增用户
}

type ChartOrder struct {
	PrePrice   *orderedmap.OrderedMap `json:"prePrice"`   // 上个时间周期销售额
	PreQuality *orderedmap.OrderedMap `json:"preQuality"` // 上个时间周期订单
	Price      *orderedmap.OrderedMap `json:"price"`      // 这个时间周期销售额
	Quality    *orderedmap.OrderedMap `json:"quality"`    // 这个时间周期订单
}
