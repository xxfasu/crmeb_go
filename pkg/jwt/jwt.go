package jwt

import (
	"context"
	"crmeb_go/internal/data/user_data"
	"crmeb_go/pkg/cache"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	cache *cache.Cache
}

type MyCustomClaims struct {
	UserID string
	jwt.RegisteredClaims
}

const (
	MillisMinuteTen = 20 * 60
	MillisMinute    = 60 * time.Minute
	ExpireTime      = 5 * 60 * 60
	RedisTokenKey   = "TOKEN:ADMIN:"
)

func NewJwt(cache *cache.Cache) *JWT {
	return &JWT{cache}
}

func (j *JWT) GenToken(userID string, expiresAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "",
			Subject:   "",
			ID:        "",
			Audience:  []string{},
		},
	})

	// Sign and get the complete encoded token as a string using the key
	tokenString, err := token.SignedString("security.jwt.key")
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) ParseToken(tokenString string) (*MyCustomClaims, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	if strings.TrimSpace(tokenString) == "" {
		return nil, errors.New("token is empty")
	}
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (j *JWT) CreateToken(loginUserData user_data.LoginUserData) (string, error) {
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

func (j *JWT) GetLoginUser(ctx *gin.Context) (user_data.LoginUserData, error) {
	var loginUserData user_data.LoginUserData
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

func (j *JWT) VerifyToken(loginUserData user_data.LoginUserData) error {
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

func (j *JWT) RefreshToken(loginUserData user_data.LoginUserData) error {
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
