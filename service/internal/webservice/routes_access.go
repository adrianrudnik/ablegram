package webservice

import (
	"github.com/adrianrudnik/ablegram/crypt"
	"github.com/adrianrudnik/ablegram/internal/access"
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/gin-gonic/gin"
)

func registerAccessRoutes(rg *gin.RouterGroup, conf *config.Config, auth *access.Auth, otp *access.Otp) {
	// Used to identify the current session against the server
	rg.GET("/auth/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"username": c.MustGet("username").(string),
			"ip":       c.ClientIP(),
			"role":     c.MustGet("role").(string),
		})
	})

	rg.POST("/auth/goodbye", func(c *gin.Context) {
		c.SetCookie("ablegram-token", "", -1, "/", "", false, true)
		c.JSON(200, gin.H{})
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

		t, err := auth.ConvertOtpToToken(input.OtpToken)
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

	rg.POST("/auth/password", func(c *gin.Context) {
		type userInput struct {
			Password string `json:"password"`
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

		setAdminCookie(c, auth)

		c.JSON(200, gin.H{})
	})
}

func setAdminCookie(c *gin.Context, auth *access.Auth) {
	token, err := auth.CreateToken()
	if err != nil {
		Logger.Error().Err(err).Msg("Could not create token")
		return
	}

	c.SetCookie("ablegram-token", token, 0, "/", "", false, true)
}
