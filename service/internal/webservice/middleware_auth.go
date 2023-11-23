package webservice

import (
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/access"
	"github.com/gin-gonic/gin"
	"strings"
)

func AccessMiddleware(auth *access.Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		displayName := "Guest" // used for UI display
		role := "guest"        // used for authorization

		// Resolve admin tokens
		token, err := c.Cookie("ablegram-token")
		if err == nil && auth.ValidateToken(token) {
			role = "admin"
			displayName = "Admin"
		}

		// Resolve a custom username
		username, err := c.Cookie("ablegram-username")
		if err == nil {
			displayName = username
		}

		// Decide on a clean and sanitized display name
		switch role {
		case "admin":
			// Admins get their IP stripped
			displayName = fmt.Sprintf("%.16s", strings.TrimSpace(displayName))
		default:
			// Everyone else gets their IP appended
			displayName = fmt.Sprintf("%.16s [%s]", strings.TrimSpace(displayName), c.ClientIP())
		}

		c.Set("role", role)
		c.Set("username", displayName)
	}
}
