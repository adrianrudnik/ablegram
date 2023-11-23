package webservice

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func isAdmin(c *gin.Context) bool {
	if c.GetString("role") != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		c.Abort()

		return false
	}

	return true
}
