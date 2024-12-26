package middleware

import (
	"crmeb_go/internal/casbin"
	"crmeb_go/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CasbinM struct {
	casbin casbin.Service
	jwt    *jwt.JWT
}

func NewCasbinM(casbin casbin.Service, jwt *jwt.JWT) *CasbinM {
	return &CasbinM{
		casbin: casbin,
		jwt:    jwt,
	}
}

func (m *CasbinM) CasbinMiddleware(obj string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文或JWT中提取用户角色
		loginUser, err := m.jwt.GetLoginUser(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No user found"})
			return
		}
		if loginUser.User == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No user found"})
			return
		}
		userRole := strconv.Itoa(int(loginUser.User.ID))
		if userRole == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No role found"})
			return
		}

		ok, err := m.casbin.Enforce(userRole, obj, "ALL")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			return
		}

		c.Next()
	}
}
