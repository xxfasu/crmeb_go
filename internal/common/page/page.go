package page

import "crmeb_go/internal/validation"

type CommonPageResp[T any] struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalPage int `json:"total_page"`
	Total     int `json:"total"`
	List      []T `json:"list"`
}

func RestPage[T any](param validation.PageParam, list []T, total int64) *CommonPageResp[T] {
	p := new(CommonPageResp[T])
	p.Page = param.Page
	p.Limit = param.Limit
	p.Total = int(total)
	p.List = list
	p.TotalPage = p.Total / p.Limit
	if p.Total%p.Limit != 0 {
		p.TotalPage++
	}
	return p
}
