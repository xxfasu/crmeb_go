package validation

type PageParam struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}
