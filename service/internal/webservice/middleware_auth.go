package webservice

import (
	"github.com/adrianrudnik/ablegram/internal/access"
	"github.com/adrianrudnik/ablegram/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccessMiddleware(auth *access.Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		displayName := "Guest"   // used for UI display
		role := access.GuestRole // used for authorization

		// Resolve admin tokens
		token, err := c.Cookie("ablegram-token")
		if err == nil && auth.ValidateToken(token) {
			role = access.AdminRole
			displayName = "Admin"
		}

		// Resolve a custom username
		username, err := c.Cookie("ablegram-username")
		if err == nil {
			displayName = username
		}

		c.Set("displayName", util.SanitizeDisplayName(displayName))
		c.Set("role", role)
	}
}

func isAdmin(c *gin.Context) bool {
	if c.GetString("role") != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		c.Abort()

		return false
	}

	return true
}

func isGuest(c *gin.Context) bool {
	if c.GetString("role") != "guest" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		c.Abort()

		return false
	}

	return true
}

func isSomeone(c *gin.Context) bool {
	if c.GetString("role") == "" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		c.Abort()

		return false
	}

	return true
}
