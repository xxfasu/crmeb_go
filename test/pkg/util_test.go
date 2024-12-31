package pkg

import (
	"crmeb_go/pkg/util"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestBcrypt(t *testing.T) {
	convey.Convey("测试 HashPassword 和 ComparePasswords 功能", t, func() {
		type TestCase struct {
			name           string
			password       string
			expectedResult bool
		}

		testCases := []TestCase{
			{
				name:     "Valid Password",
				password: "123456",
			},
			{
				name:     "Invalid Password",
				password: "654321",
			},
			{
				name:     "Empty Password",
				password: "",
			},
		}

		for _, tc := range testCases {
			convey.Convey(tc.name, func() {
				// 哈希密码
				hashedPassword, err := util.HashPassword(tc.password)
				convey.So(err, convey.ShouldBeNil)
				convey.So(hashedPassword, convey.ShouldNotBeEmpty)

				// 比较密码（正确密码）
				result := util.ComparePasswords(hashedPassword, tc.password)
				convey.So(result, convey.ShouldBeTrue)

				// 比较密码（错误密码）
				result = util.ComparePasswords(hashedPassword, "wrongpassword")
				convey.So(result, convey.ShouldBeFalse)
			})
		}
	})
}
