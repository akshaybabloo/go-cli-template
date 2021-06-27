package configuration

import (
	"os"
	"path"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

type GlobalConfig struct {
	EnableColour bool   `yaml:"enable_colour"`
	ID           string `yaml:"id"`
}

type Configuration interface {
	Base() (string, error)
	StatePath() (string, error)
	GlobalConfig() (GlobalConfig, error)
	GetGlobalConfigPath() (string, error)
}

func NewConfig() Configuration {
	return &configStubs{}
}

type configStubs struct {
	Configuration
}

// Base path for config
func (c *configStubs) Base() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	rexFolder := path.Join(home, ".config", "rex")
	return rexFolder, nil
}

// StatePath returns the state path
func (c *configStubs) StatePath() (string, error) {
	home, err := c.Base()
	if err != nil {
		return "", err
	}
	return path.Join(home, "state.yaml"), nil
}

// GetGlobalConfigPath returns the global config path
func (c *configStubs) GetGlobalConfigPath() (string, error) {
	home, err := c.Base()
	if err != nil {
		return "", err
	}
	return path.Join(home, "config.yaml"), nil
}

// GlobalConfig return the global configuration
func (c *configStubs) GlobalConfig() (GlobalConfig, error) {
	var globalConfig GlobalConfig

	configPath, err := c.GetGlobalConfigPath()
	if err != nil {
		return GlobalConfig{}, err
	}

	err = unmarshalProfile(configPath, &globalConfig)
	if err != nil {
		return GlobalConfig{}, err
	}

	return globalConfig, nil
}

func unmarshalProfile(path string, o interface{}) error {
	fs := afero.NewOsFs()
	appFs := afero.Afero{Fs: fs}

	file, err := appFs.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, o)
	if err != nil {
		return err
	}

	return nil
}
