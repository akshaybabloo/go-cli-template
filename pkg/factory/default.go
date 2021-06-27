package factory

import (
	"github.com/akshaybabloo/go-cli-template/pkg/colours"
	"github.com/akshaybabloo/go-cli-template/pkg/configuration"
)

type Factory struct {
	Config  func() configuration.Configuration
	Debug   bool
	Colours func() colours.Colours
}

func New() *Factory {
	return &Factory{
		Config: configuration.NewConfig,
		Debug:  false,
		Colours: func() colours.Colours {
			return colours.NewColour(configuration.NewConfig())
		},
	}
}
