package util

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"log"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBcrypt(t *testing.T) {
	Convey("测试 HashPassword 和 ComparePasswords 功能", t, func() {
		type TestCase struct {
			name            string
			password        string
			comparePassword string
			assert          Assertion
		}

		testCases := []TestCase{
			{
				name:            "校验正确的密码",
				password:        "123456",
				comparePassword: "123456",
				assert:          ShouldBeTrue,
			},
			{
				name:            "校验错误的密码",
				password:        "123456",
				comparePassword: "654321",
				assert:          ShouldBeFalse,
			},
			{
				name:            "校验空密码",
				password:        "123456",
				comparePassword: "",
				assert:          ShouldBeFalse,
			},
		}

		for _, tc := range testCases {
			Convey(tc.name, func() {
				// 哈希密码
				hashedPassword, err := HashPassword(tc.password)
				So(err, ShouldBeNil)
				So(hashedPassword, ShouldNotBeEmpty)

				// 比较密码
				result := ComparePasswords(hashedPassword, tc.comparePassword)
				So(result, tc.assert)
			})
		}
	})
}

func TestCopierOption(t *testing.T) {
	Convey("测试 Copier.CopyWithOption 方法", t, func() {
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

		Convey("执行 Copier.CopyWithOption 进行数据复制", func() {
			// 执行复制操作
			err := copier.CopyWithOption(m, u, convertOption(*m, *u))

			// 断言没有错误发生
			So(err, ShouldBeNil)
			// 断言字段是否正确复制和转换
			Printf("结果为:%#v", *m)
			Convey("验证字段转换是否正确", func() {
				So(m.Status, ShouldBeTrue)                            // 假设 Status:1 转换为 true
				So(m.CreatedTime, ShouldEqual, "2024-01-10 08:00:00") // 假设 CreatedAt:111111111 转换为字符串 "111111111"
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

func TestGetColumns(t *testing.T) {
	// 1. 连接数据库
	dsn := "root:password@tcp(127.0.0.1:3306)/crmeb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	defer db.Close()

	// 2. 获取指定表的所有列名
	tableName := "eb_system_admin" // 假设要获取 eb_system_admin 表的字段
	cols, err := getColumns(db, tableName)
	if err != nil {
		log.Fatalf("获取表 %s 字段信息失败: %v", tableName, err)
	}

	fmt.Println(formatColumns(cols, 1))
}
