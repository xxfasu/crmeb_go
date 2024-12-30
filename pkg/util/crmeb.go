package util

import (
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

func StrToArrInt64(str string) []int64 {
	strArr := strings.Split(str, ",")
	arr := make([]int64, 0, len(strArr))
	for _, s := range strArr {
		i, _ := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
		arr = append(arr, i)
	}
	return arr
}

// HashPassword 使用 Bcrypt 算法生成密码哈希值
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePasswords 比较输入的密码与哈希值是否匹配
func ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
