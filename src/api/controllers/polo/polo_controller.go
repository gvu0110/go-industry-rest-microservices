package polo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	polo = "polo"
)

// Macro function to make sure the webapp is working
func Macro(c *gin.Context) {
	c.String(http.StatusOK, polo)
}
