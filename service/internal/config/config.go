package config

import (
	"errors"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

func newConfig() *Config {
	return &Config{
		Log: LogConfig{
			Level:                  "info",
			EnableRuntimeLogfile:   false,
			EnableProcessedLogfile: false,
			RuntimeLogfilePath:     GetRelativeFilePath(".runtime.log"),
			ProcessLogfilePath:     GetRelativeFilePath(".processed.log"),
		},

		Behaviour: BehaviourConfig{
			WebserviceAutostart: true,
			BrowserAutostart:    true,
			ShowGui:             true,
		},

		Collector: CollectorConfig{
			SearchablePaths:      make([]string, 0, 100),
			ExcludeSystemFolders: true,
			WorkerCount:          5,
			WorkerDelayInMs:      0,
		},

		Parser: ParserConfig{
			WorkerCount:     5,
			WorkerDelayInMs: 0,
		},

		Webservice: WebserviceConfig{
			TryPorts: []int{10000, 20000, 30000, 40000, 50000, 10001},
		},
	}
}

// LoadWithDefaults tries to load an existing configuration, but falls back to defaults if none could be found or read.
// If path is an empty string, a config file with the name of the executable  with an ".config.yaml" extension will be tried.
func LoadWithDefaults(path string) *Config {
	c, err := Load(path)
	if err != nil {
		// Create fallback configuration
		c := newConfig()

		// For the searchable paths, we prefer the users home directory as initial configuration
		homeDir, err := os.UserHomeDir()
		if err == nil {
			c.Collector.SearchablePaths = append(c.Collector.SearchablePaths, homeDir)
		}

		return c
	}

	return c
}

// Load tries to load an existing configuration near the executable.
// If path is an empty string, a config file with the name of the executable with an ".config.yaml" extension will be tried.
func Load(path string) (*Config, error) {
	if path == "" {
		path = GetRelativeFilePath(".config.yaml")
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := newConfig()

	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Config) Save() error {
	b, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}

	err = os.WriteFile(GetRelativeFilePath(".config.yaml"), b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetRelativeFilePath(sub string) string {
	absPath, err := func() (string, error) {
		p, err := os.Executable()
		if err != nil {
			return "", err
		}

		if strings.Contains(p, ".cache") && strings.Contains(p, "GoLand") {
			Logger.Warn().Err(err).Msg("Detected developer IDE, falling back to working directory for for executable path")
			return "", errors.New("developer")
		}

		return p, err
	}()

	if err != nil {
		Logger.Warn().Err(err).Msg("Failed to find the path of the executable, falling back to working directory")
		absPath, err = os.Getwd()
		if err != nil {
			Logger.Warn().Err(err).Msg("Failed to find working directory, falling back to simple ./")
			absPath = "./Ablegram"
		} else {
			absPath += "/Ablegram"
		}
	}

	// Remove a possible extension, like with win
	base := strings.TrimSuffix(absPath, filepath.Ext(absPath))
	base += sub

	return base
}
