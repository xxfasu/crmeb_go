package middleware

import (
	"crmeb_go/pkg/jwt"
	"crmeb_go/pkg/logs"
	"crmeb_go/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthM struct {
	j *jwt.JWT
}

func NewAuthM(j *jwt.JWT) *AuthM {
	return &AuthM{

		j: j,
	}
}

func (m *AuthM) StrictAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// tokenString := ctx.Request.Header.Get("Authorization")
		loginUserData, err := m.j.GetLoginUser(ctx)
		if err != nil {
			logs.Log.WithContext(ctx).Warn("No token", zap.Any("data", map[string]interface{}{
				"url":    ctx.Request.URL,
				"params": ctx.Params,
			}))
			response.FailWithMessage(ctx, "token is empty")
			ctx.Abort()
			return
		}

		ctx.Set("user-info", loginUserData)
		ctx.Next()
	}
}

func (m *AuthM) NoStrictAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loginUserData, err := m.j.GetLoginUser(ctx)
		if err != nil {
			ctx.Next()
			return
		}

		ctx.Set("resp", loginUserData)
		ctx.Next()
	}
}
