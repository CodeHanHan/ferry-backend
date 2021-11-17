package apis

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlePing(c *gin.Context) {
	message := c.Query("message")

	reply := fmt.Sprintf("%s, too", message)

	c.String(http.StatusOK, reply)
}
