package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/CodeHanHan/ferry-backend/pkg/jwtauth"
)

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type User struct {
	UserName string
	Email    string
}

var identityKey = "id"

// Login godoc
// @Summary 用户登录，获取token
// @Description 账号密码登录(测试)
// @Tags user
// @ID user-login
// @Param username query string true "any string"
// @Param password query string true "any string"
// @Success 200 {string} string "{"code": 200, "expire": "2019-08-07T12:45:48+08:00", "token": ".eyJleHAiOjE1NjUxNTMxNDgsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU2NTE0OTU0OH0.-zvzHvbg0A" }"
// @Accept  application/json
// @Produce  application/json
// @Router /user/login [post]
func Authenticator(c *gin.Context) (interface{}, error) {
	var loginVals Login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwtauth.ErrMissingLoginValues
	}

	userId := loginVals.Username
	password := loginVals.Password

	if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
		return &User{
			UserName: userId,
			Email:    fmt.Sprintf("%s@ferry.com", userId),
		}, nil
	}

	return nil, jwtauth.ErrFailedAuthentication
}

// 应该对经过身份验证的用户执行授权的回调函数。 仅在身份验证成功后调用。 成功时必须返回真，失败时必须返回假。
func Authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*User); ok && v.UserName == "admin" {
		return true
	}

	return false
}

// 登录时调用的回调函数。
// 使用此功能可以向网络令牌添加额外的有效载荷数据。
// 然后在请求期间通过 c.Get("JWT_PAYLOAD") 访问数据。
// 请注意，有效负载未加密。
// jwt.io 上提到的属性不能用作地图的键。
// 可选，默认情况下不会设置额外的数据。
func PayloadFunc(data interface{}) jwtauth.MapClaims {
	if v, ok := data.(*User); ok {
		return jwtauth.MapClaims{
			jwtauth.IdentityKey: v.UserName,
		}
	}
	return jwtauth.MapClaims{}
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwtauth.ExtractClaims(c)
	return &User{
		UserName: claims[identityKey].(string),
	}
}
