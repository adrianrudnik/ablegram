package ui

import (
	"github.com/icza/gox/osx"
)

func OpenFrontend() {
	err := osx.OpenDefault("http://localhost:10000")
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
