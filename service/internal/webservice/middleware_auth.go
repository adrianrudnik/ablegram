package webservice

import (
	"github.com/adrianrudnik/ablegram/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func AuthMiddleware(auth *auth.Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Try to read the raw token, if present
		rawToken, err := getAuthCookie(c)
		if err != nil {
			// No token found, we are done here
			setAuthContext(nil, c)
			return
		}

		token, ok := auth.ValidateToken(rawToken)
		if !ok {
			// Invalid token, we are done here
			setAuthContext(nil, c)
			return
		}

		setAuthContext(token, c)
	}
}

// setAuthCookie initializes the required context variables for all available routes.
func setAuthContext(token *auth.AuthToken, c *gin.Context) {
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

// isAdmin checks if there is a known user and if he has the admin role.
// If not, the given context will be aborted with a 403 Forbidden and false will be returned.
func isAdmin(c *gin.Context) bool {
	if c.GetBool("user") && c.GetString("userRole") != auth.AdminRole {
		c.AbortWithStatus(http.StatusForbidden)
		return false
	}

	return true
}

// isSomeone checks if there is a known user.
// If not, the given context will be aborted with a 403 Forbidden and false will be returned.
func isSomeone(c *gin.Context) bool {
	if !c.GetBool("user") {
		c.AbortWithStatus(http.StatusForbidden)

		return false
	}

	return true
}
