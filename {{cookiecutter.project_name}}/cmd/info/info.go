package info

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/akshaybabloo/go-cli-template/pkg/factory"
)

type Info struct {
	Version            string
	InstalledGoVersion string
	GoVersionBuildWith string
	GitVersion         string
	Platform           string
	DockerVersion      string
	Architecture       string
}

func NewInfoCmd(f *factory.Factory, v string) *cobra.Command {
	var infoCmd = &cobra.Command{
		Use:   "info",
		Short: "Information of rex and it's variables",
		Run: func(cmd *cobra.Command, args []string) {
			path, err := f.Config().GetGlobalConfigPath()
			if err != nil {
				panic(err)
			}
			fmt.Println("Global configuration path", path)
			fmt.Println("Version", v)
		},
	}

	return infoCmd
}
