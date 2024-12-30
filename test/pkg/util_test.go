package pkg

import (
	"crmeb_go/pkg/util"
	"fmt"
	"testing"
)

func TestBcrypt(t *testing.T) {
	password, err := util.HashPassword("123456")
	if err != nil {
		return
	}
	fmt.Println(password)
	fmt.Println(util.ComparePasswords(password, "123456"))
}
