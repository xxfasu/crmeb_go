package user_repository

import "crmeb_go/internal/common/data"

type Querier interface {
	//	SELECT
	//	 DATE(FROM_UNIXTIME(created_at)) AS every_date,
	//   COUNT(id) AS id,
	//  FROM
	//		eb_user
	// 	{{where}}
	// 		{{if condition.Start !=0}}
	// 			eb_user.created_at >= @condition.Start AND
	// 		{{end}}
	// 		{{if condition.End !=0}}
	// 			eb_user.created_at <  @condition.End AND
	// 		{{end}}
	// 		eb_user.deleted_at = 0
	// 	{{end}}
	// GROUP BY every_date
	// ORDER BY every_date ASC
	GetAddUserCountGroupDate(condition *data.DateCondition) ([]*data.UserEveryDate, error)
}
