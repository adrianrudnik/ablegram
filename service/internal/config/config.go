package config

import (
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

type Config struct {
	PrimaryLanguage string   `yaml:"primary_search_language"`
	SearchablePaths []string `yaml:"searchable_paths"`
}

func newConfig() *Config {
	return &Config{
		PrimaryLanguage: "en",
		SearchablePaths: make([]string, 0, 100),
	}
}

// LoadWithDefaults tries to load an existing configuration, but falls back to defaults if none could be found or read.
// If path is an empty string, a config file with the name of the executable  with an ".config.yaml" extension will be tried.
func LoadWithDefaults(path string) *Config {
	c, err := Load(path)
	if err != nil {
		// Create fallback configuration
		c := newConfig()

		Logger.Info().Err(err).Msg("Could not load config, falling back to defaults")

		c.PrimaryLanguage = ""

		// For the searchable paths, we prefer the users home directory as initial configuration
		homeDir, err := os.UserHomeDir()
		if err == nil {
			c.SearchablePaths = append(c.SearchablePaths, homeDir)
		}

		return c
	}

	return c
}

// Load tries to load an existing configuration near the executable.
// If path is an empty string, a config file with the name of the executable with an ".config.yaml" extension will be tried.
func Load(path string) (*Config, error) {
	if path == "" {
		p, err := getDefaultPath()
		if err != nil {
			return nil, err
		}

		path = p
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

func Save(c *Config, path string) error {
	b, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func getDefaultPath() (string, error) {
	execAbsPath, err := os.Executable()
	if err != nil {
		return "", err
	}

	// Remove a possible extension, like with win
	base := strings.TrimSuffix(execAbsPath, filepath.Ext(execAbsPath))

	base += ".config.yaml"

	return base, nil
}
