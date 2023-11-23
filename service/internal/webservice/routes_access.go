package webservice

import (
	"github.com/adrianrudnik/ablegram/internal/access"
	"github.com/gin-gonic/gin"
)

func registerAccessRoutes(rg *gin.RouterGroup, auth *access.Auth, otp *access.Otp) {
	// Used to identify the current session against the server
	rg.GET("/auth/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"username": c.MustGet("username").(string),
			"role":     c.MustGet("role").(string),
		})
	})

	rg.POST("/auth/otp", func(c *gin.Context) {
		type userInput struct {
			OtpToken string `json:"token"`
		}

		var input userInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		t, err := auth.ConvertOtp(input.OtpToken)
		if err != nil {
			Logger.Warn().
				Str("ip", c.ClientIP()).
				Msg("Tried to authenticate with invalid OTP token")

			c.JSON(403, gin.H{"error": err.Error()})
			return
		}

		Logger.Info().
			Str("ip", c.ClientIP()).
			Msg("Authenticated with valid OTP token")

		c.SetCookie("ablegram-token", t, 0, "/", "", false, true)

		c.JSON(200, gin.H{})
	})
}
