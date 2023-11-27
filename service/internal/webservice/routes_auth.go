package webservice

import (
	"fmt"
	"github.com/adrianrudnik/ablegram/crypt"
	"github.com/adrianrudnik/ablegram/internal/access"
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func registerAuthRoutes(rg *gin.RouterGroup, conf *config.Config, auth *access.Auth) {
	// Every client needs to say hello.
	// The service will issue a cookie that the client needs to serve on every other request.
	rg.POST("/auth", func(c *gin.Context) {
		// Users that we identify through a token will receive their info back and some.
		if c.GetBool("user") {

			c.JSON(200, access.User{
				UserID:      c.MustGet("userId").(uuid.UUID),
				DisplayName: c.GetString("userDisplayName"),
				Role:        c.GetString("userRole"),
			})

			return
		}

		// Users that we could not identify will receive a token that they need to serve on every other request.
		token := access.NewGuestAuthToken()

		setAuthCookie(c, token)

		c.JSON(200, access.User{
			UserID:      token.ID,
			DisplayName: token.DisplayName,
			Role:        token.Role,
		})
	})

	// Allows a current user to set a custom display name
	rg.PUT("/auth/display-name", func(c *gin.Context) {
		if !isSomeone(c) {
			return
		}

		type userInput struct {
			DisplayName string `json:"display_name" binding:"required,alphanumunicode,min=3,max=16"`
		}

		var input userInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Get the current token and update the display name
		t := c.MustGet("userToken").(*access.AuthToken)
		t.DisplayName = input.DisplayName

		setAuthCookie(c, t)

		c.JSON(200, gin.H{
			"display_name": t.DisplayName,
		})
	})

	rg.POST("/auth/otp", func(c *gin.Context) {
		type userInput struct {
			OtpToken string `json:"token" binding:"required"`
		}

		var input userInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		t, err := auth.ConvertOtpToAdminToken(input.OtpToken)
		if err != nil {
			Logger.Warn().
				Str("ip", c.ClientIP()).
				Msg("Tried to authenticate with invalid OTP token")

			c.JSON(403, gin.H{"error": err.Error()})
			return
		}

		Logger.Info().
			Str("user", t.ID.String()).
			Str("ip", c.ClientIP()).
			Msg("Authenticated with valid OTP token")

		setAuthCookie(c, t)

		c.JSON(200, gin.H{})
	})

	rg.POST("/auth/password", func(c *gin.Context) {
		type userInput struct {
			Password string `json:"password" binding:"required"`
		}

		var input userInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if conf.Webservice.MasterPassword == "" {
			c.JSON(405, gin.H{"error": "Master password not enabled"})
			return
		}

		pw, err := crypt.Decrypt(conf.Webservice.MasterPassword)
		if err != nil {
			c.JSON(500, gin.H{"error": "Could not decrypt master password"})
			return
		}

		if input.Password != pw {
			c.JSON(403, gin.H{"error": "Invalid password"})
			return
		}

		t := access.NewAdminToken()

		setAuthCookie(c, t)

		c.JSON(200, gin.H{})
	})

	// Allow anyone to delete the current session cookie, we do not care if a valid session exists.
	rg.DELETE("/auth", func(c *gin.Context) {
		c.SetCookie(getCookieName(), "", -1, "/", "", false, true)
		c.JSON(200, gin.H{})
	})
}

func getCookieName() string {
	return fmt.Sprintf("ablegram-token-v%d", access.TokenVersion)
}

func setAuthCookie(c *gin.Context, token *access.AuthToken) {
	c.SetCookie(getCookieName(), token.Encrypt(), 0, "/", "", false, true)
}

func getAuthCookie(c *gin.Context) (string, error) {
	cookie, err := c.Cookie(getCookieName())
	if err != nil {
		return "", err
	}

	return cookie, nil
}
