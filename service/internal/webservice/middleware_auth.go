package webservice

import (
	"github.com/adrianrudnik/ablegram/internal/access"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func AccessMiddleware(auth *access.Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Try to read the raw token, if present
		rawToken, err := getAuthCookie(c)
		if err != nil {
			// No token found, we are done here
			setContext(nil, c)
			return
		}

		token, ok := auth.ValidateToken(rawToken)
		if !ok {
			// Invalid token, we are done here
			setContext(nil, c)
			return
		}

		setContext(token, c)
	}
}

func setContext(token *access.AuthToken, c *gin.Context) {
	if (token == nil) || (token.ID == uuid.Nil) {
		c.Set("user", false)
		return
	}

	c.Set("user", true)
	c.Set("userId", token.ID)
	c.Set("userRole", token.Role)
	c.Set("userDisplayName", token.DisplayName)
	c.Set("userToken", token)
}

func isAdmin(c *gin.Context) bool {
	if c.GetString("userRole") != access.AdminRole {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		c.Abort()

		return false
	}

	return true
}

func isSomeone(c *gin.Context) bool {
	if !c.GetBool("user") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		c.Abort()

		return false
	}

	return true
}
