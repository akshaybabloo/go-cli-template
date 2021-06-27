package model

import (
	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
)

type GlobalConfig struct {
	EnableColour bool   `yaml:"enable_colour"`
	ID           string `yaml:"id"`
}

// New creates a default value for global config
func New() GlobalConfig {
	u := uuid.New()

	return GlobalConfig{
		EnableColour: true,
		ID:           u.String(),
	}
}

// MarshalNew marshals the default struct
func MarshalNew() ([]byte, error) {
	defaultData := New()
	marshal, err := yaml.Marshal(&defaultData)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}
