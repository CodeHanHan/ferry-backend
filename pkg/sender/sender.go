package sender

import (
	"errors"

	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

var (
	ErrSenderNotExist = errors.New("get sender failed")
)

func GetSender(c *gin.Context) (username, role string, err error) {
	usernameI, ok1 := c.Get("sender")
	roleI, ok2 := c.Get("role")
	if !(ok1 && ok2) {
		logger.Error(c, ErrSenderNotExist.Error())
		return "", "", ErrSenderNotExist
	}

	username = usernameI.(string)
	role = roleI.(string)

	return
}
