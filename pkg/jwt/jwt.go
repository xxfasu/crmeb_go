package jwt

import (
	"context"
	"crmeb_go/internal/common/data/login_user"
	"crmeb_go/pkg/cache"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type JWT struct {
	cache cache.Cache
}

const (
	MillisMinuteTen = 20 * 60
	ExpireTime      = 5 * 60 * 60
	RedisTokenKey   = "TOKEN:ADMIN:"
)

func NewJwt(cache cache.Cache) *JWT {
	return &JWT{cache}
}

func (j *JWT) CreateToken(loginUserData login_user.LoginUserData) (string, error) {
	token := strings.Replace(loginUserData.Token, "-", "", -1)
	loginUserData.Token = token
	err := j.RefreshToken(loginUserData)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (j *JWT) DeleteToken(token string) error {
	key := getTokenKey(token)
	return j.cache.DelCache(key)
}

func (j *JWT) GetLoginUser(ctx *gin.Context) (login_user.LoginUserData, error) {
	var loginUserData login_user.LoginUserData
	token := getToken(ctx)
	if len(token) != 0 {
		key := getTokenKey(token)
		value, err := j.cache.GetCache(context.Background(), key)
		if err != nil {
			return loginUserData, err
		}
		err = json.Unmarshal([]byte(value), &loginUserData)
		if err != nil {
			return loginUserData, err
		}
	}

	return loginUserData, nil
}

func (j *JWT) VerifyToken(loginUserData login_user.LoginUserData) error {
	expireTime := loginUserData.ExpireTime
	currentTime := time.Now().Unix() // 获取当前时间，Unix 时间戳，单位为秒
	if expireTime-currentTime <= MillisMinuteTen {
		err := j.RefreshToken(loginUserData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (j *JWT) RefreshToken(loginUserData login_user.LoginUserData) error {
	loginUserData.LoginTime = time.Now().Unix()
	loginUserData.ExpireTime = loginUserData.LoginTime + ExpireTime
	key := getTokenKey(loginUserData.Token)
	marshal, _ := json.Marshal(loginUserData)
	err := j.cache.SetCache(context.Background(), key, string(marshal), time.Duration(loginUserData.ExpireTime)*time.Second)
	if err != nil {
		return err
	}
	return nil
}

func getTokenKey(uuid string) string {
	return RedisTokenKey + uuid
}

func getToken(ctx *gin.Context) string {
	token := ctx.Request.Header.Get("Authori-zation")
	if token != "" {
		token = strings.Replace(token, RedisTokenKey, "", -1)
	}
	return token
}
