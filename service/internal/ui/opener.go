package ui

import (
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/access"
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/icza/gox/osx"
	"net/url"
)

func GenerateLocalAdminUrl(c *config.Config, otp *access.Otp) string {
	port := c.Webservice.ChosenPort

	if c.IsDevelopmentEnv {
		port = 5173
	}

	otpToken := otp.CreateOtp()
	return fmt.Sprintf("http://localhost:%d/auth/otp?token=%s", port, url.QueryEscape(otpToken))
}

func OpenFrontendAsAdmin(c *config.Config, otp *access.Otp) {
	err := osx.OpenDefault(GenerateLocalAdminUrl(c, otp))
	if err != nil {
		Logger.Warn().Err(err).Msg("Could not open default browser")
	}
}

func OpenDefault(path string) {
	err := osx.OpenDefault(path)

	if err != nil {
		Logger.Warn().Err(err).Msg("Could not open default browser")
	}
}
