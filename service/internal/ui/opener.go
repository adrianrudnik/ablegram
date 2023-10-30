package ui

import (
	"fmt"
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/icza/gox/osx"
)

func OpenFrontend(c *config.Config) {
	err := osx.OpenDefault(fmt.Sprintf("http://localhost:%d", c.Webservice.ChosenPort))
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
