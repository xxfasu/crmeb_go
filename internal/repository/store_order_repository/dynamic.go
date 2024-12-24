package store_order_repository

import "crmeb_go/internal/common/data"

type Querier interface {
	//	SELECT
	//	 DATE(FROM_UNIXTIME(created_at)) AS every_date,
	//   COUNT(id) AS id,
	//   SUM(pay_price) AS pay_price
	//  FROM
	//		eb_store_order
	// 	{{where}}
	// 		{{if condition.Start !=0}}
	// 			eb_store_order.created_at >= @condition.Start AND
	// 		{{end}}
	// 		{{if condition.End !=0}}
	// 			eb_store_order.created_at <  @condition.End AND
	// 		{{end}}
	// 		eb_store_order.deleted_at = 0
	// 	{{end}}
	// GROUP BY every_date
	// ORDER BY every_date ASC
	QueryOrderGroupByDate(condition *data.DateCondition) ([]*data.StoreOrderEveryDate, error)
}
