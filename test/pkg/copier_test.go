package pkg

import (
	"errors"
	"github.com/jinzhu/copier"
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCopierOption(t *testing.T) {
	convey.Convey("测试 Copier.CopyWithOption 方法", t, func() {
		// 定义源结构体 User1
		type User1 struct {
			Status    int64 `json:"status"`
			CreatedAt int64 `json:"created_at"`
		}

		// 定义目标结构体 User2
		type User2 struct {
			Status      bool   `json:"status"`
			CreatedTime string `json:"created_time"`
		}

		// 初始化源数据
		u := &User1{
			Status:    1,
			CreatedAt: 1704844800,
		}

		// 初始化目标数据
		m := new(User2)

		convey.Convey("执行 Copier.CopyWithOption 进行数据复制", func() {
			// 执行复制操作
			err := copier.CopyWithOption(m, u, convertOption(*m, *u))

			// 断言没有错误发生
			convey.So(err, convey.ShouldBeNil)
			// 断言字段是否正确复制和转换
			convey.Convey("验证字段转换是否正确", func() {
				convey.So(m.Status, convey.ShouldBeFalse)                           // 假设 Status:1 转换为 true
				convey.So(m.CreatedTime, convey.ShouldEqual, "2024-01-10 08:00:00") // 假设 CreatedAt:111111111 转换为字符串 "111111111"
			})
		})
	})
}

func convertOption(m any, u any) copier.Option {
	return copier.Option{
		Converters: []copier.TypeConverter{
			{
				SrcType: int64(0),
				DstType: copier.String,
				Fn: func(src interface{}) (interface{}, error) {
					s, ok := src.(int64)
					if !ok {
						return nil, errors.New("src type not matching")
					}
					return time.Unix(s, 0).Format(time.DateTime), nil
				},
			},
			{
				SrcType: int64(0),
				DstType: copier.Bool,
				Fn: func(src interface{}) (interface{}, error) {

					s, ok := src.(int64)
					if !ok {
						return nil, errors.New("src type not matching")
					}
					return s == 1, nil
				},
			},
		},
		FieldNameMapping: []copier.FieldNameMapping{
			{
				SrcType: u,
				DstType: m,
				Mapping: map[string]string{
					"CreatedAt": "CreatedTime",
				}},
		}}
}
