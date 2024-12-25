package pkg

import (
	"github.com/iancoleman/orderedmap"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	// strings := []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}
	// strings := []string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"}
	strings := []string{"1-12", "2-12", "3-12", "4-12", "10-12", "11-12", "12-12"}
	orderedMap := orderedmap.New()
	for _, v := range strings {
		orderedMap.Set(v, v)
	}
	orderedMap.SortKeys(func(keys []string) {
		sort.Strings(keys)
	})
	t.Log(orderedMap.Keys())
}
