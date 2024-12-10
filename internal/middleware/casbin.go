package middleware

import (
	"crmeb_go/internal/casbin_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CasbinM struct {
	casbin casbin_service.Service
}

func NewCasbinM(casbin casbin_service.Service) *CasbinM {
	return &CasbinM{
		casbin: casbin,
	}
}

func (m *CasbinM) CasbinMiddleware(obj string, act string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文或JWT中提取用户角色
		userRole := c.GetString("user_role")
		if userRole == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No role found"})
			return
		}

		ok, err := m.casbin.Enforce(userRole, obj, act)
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
