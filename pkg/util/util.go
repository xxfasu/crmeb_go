package util

import (
	"errors"
)

// ConvertSlice 将 []any 转换为 []T，其中 T 是目标类型。
// 如果转换失败，返回错误。
func ConvertSlice[T any](input []any) ([]T, error) {
	var result []T
	for _, item := range input {
		// 使用类型断言将 item 转换为 T
		castedItem, ok := item.(T)
		if !ok {
			return nil, errors.New("转换失败")
		}
		result = append(result, castedItem)
	}
	return result, nil
}
