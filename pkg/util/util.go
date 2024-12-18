package util

import (
	"crmeb_go/internal/validation"
	"encoding/json"
)

// ConvertSlice 将 []any 转换为 []T，其中 T 是目标类型。
// 如果转换失败，返回错误。
func ConvertSlice[T any](input []any) ([]T, error) {
	var result []T
	for _, item := range input {
		// 使用类型断言将 item 转换为 T
		var temp T
		marshal, _ := json.Marshal(item)
		json.Unmarshal(marshal, &temp)
		result = append(result, temp)
	}
	return result, nil
}

func DefaultPageParams() validation.PageParam {
	return validation.PageParam{
		Page:  1,
		Limit: 20,
	}
}
