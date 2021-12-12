package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CodeHanHan/ferry-backend/pkg/token"
)

func Sender() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, err := token.GetPayload(c)
		if err != nil {
			c.AbortWithError(http.StatusForbidden, fmt.Errorf("无效的token"))
			return
		}

		c.Set("sender", payload.Username)
		c.Set("role", payload.Role)
		c.Next()
	}
}
