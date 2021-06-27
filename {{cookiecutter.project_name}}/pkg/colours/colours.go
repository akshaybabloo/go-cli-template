package colours

import (
	"strings"

	"github.com/fatih/color"

	"github.com/akshaybabloo/go-cli-template/pkg/configuration"
)

const (
	FgRed    = color.FgRed
	FgGreen  = color.FgGreen
	FgYellow = color.FgYellow
	FgCyan   = color.FgCyan
)

const (
	FgHiBlue = color.FgHiBlue
)

const (
	Bold = color.Bold
)

type Colours interface {
	New(value ...color.Attribute) *color.Color

	BoldString(format string, a ...interface{}) string
	BlueHiString(format string, a ...interface{}) string
	GreenString(format string, a ...interface{}) string
	RedString(format string, a ...interface{}) string
	YellowString(format string, a ...interface{}) string

	Yellow(format string, a ...interface{})
	Red(format string, a ...interface{})
	Green(format string, a ...interface{})

	newColour(value ...color.Attribute) *color.Color
	colorPrint(format string, p color.Attribute, a ...interface{})
}

func NewColour(conf configuration.Configuration) *colourConfig {
	return &colourConfig{config: conf}
}

type colourConfig struct {
	config configuration.Configuration
}

func (c *colourConfig) newColour(value ...color.Attribute) *color.Color {
	gConf, _ := c.config.GlobalConfig()
	co := color.New(value...)
	if !gConf.EnableColour {
		co.DisableColor()
	}
	return co
}

func (c *colourConfig) New(value ...color.Attribute) *color.Color {
	return c.newColour(value...)
}

func (c *colourConfig) colorPrint(format string, p color.Attribute, a ...interface{}) {
	co := c.newColour(p)
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	if len(a) == 0 {
		co.Print(format)
	} else {
		co.Printf(format, a...)
	}
}

func (c *colourConfig) colorString(format string, p color.Attribute, a ...interface{}) string {
	co := c.newColour(p)

	if len(a) == 0 {
		return co.SprintFunc()(format)
	}

	return co.SprintfFunc()(format, a...)
}

func (c *colourConfig) BoldString(format string, a ...interface{}) string {
	return c.colorString(format, Bold, a...)
}

func (c *colourConfig) BlueHiString(format string, a ...interface{}) string {
	return c.colorString(format, FgHiBlue, a...)
}

func (c *colourConfig) GreenString(format string, a ...interface{}) string {
	return c.colorString(format, FgGreen, a...)
}

func (c *colourConfig) RedString(format string, a ...interface{}) string {
	return c.colorString(format, FgRed, a...)
}

func (c *colourConfig) YellowString(format string, a ...interface{}) string {
	return c.colorString(format, FgYellow, a...)
}

func (c *colourConfig) Yellow(format string, a ...interface{}) {
	c.colorPrint(format, FgYellow, a...)
}

func (c *colourConfig) Red(format string, a ...interface{}) {
	c.colorPrint(format, FgRed, a...)
}

func (c *colourConfig) Green(format string, a ...interface{}) {
	c.colorPrint(format, FgGreen, a...)
}
