package user_repository

import (
	"crmeb_go/internal/data/user_data"
	"gorm.io/gen"
)

type Querier interface {
	//	SELECT id,user_id,nickname,email FROM users
	// 		{{where}}
	// 			{{if condition.Nickname !=""}}
	// 				nickname = @condition.Nickname AND
	// 			{{end}}
	// 			{{if condition.UserID !=""}}
	// 				user_id = @condition.UserID AND
	// 			{{end}}
	// 			{{if condition.Email !=""}}
	// 				email = @condition.Email
	// 			{{end}}
	// 		{{end}}
	GetUserByCondition(condition user_data.Condition) (*gen.T, error)
}
