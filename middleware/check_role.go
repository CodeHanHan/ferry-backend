package middleware

import (
	"fmt"
	"net/http"

	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/pkg/pi"
	"github.com/CodeHanHan/ferry-backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func CheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, err := token.GetPayload(c)
		if err != nil {
			c.AbortWithError(http.StatusForbidden, fmt.Errorf("无效的token"))
			return
		}

		sub := payload.Role
		obj := c.Request.URL.Path
		act := c.Request.Method
		logger.Info(c, sub, obj, act)

		ok, err := pi.Global.Casbin.Enforce(sub, obj, act)
		if err != nil {
			logger.Error(c, "鉴权错误")
			return
		}

		if ok {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  fmt.Sprintf("对不起，您没有 <%v-%v> 访问权限，请联系管理员", obj, act),
			})
			c.Abort()
			return
		}
	}
}
