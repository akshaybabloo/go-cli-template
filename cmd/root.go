package cmd

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/akshaybabloo/go-cli-template/cmd/version"
	"github.com/akshaybabloo/go-cli-template/pkg/factory"
)

func NewRootCmd(f *factory.Factory, appVersion, buildDate string) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "rex [OPTIONS] [COMMANDS]",
		Short: "Release exporter",
		Long:  `rex is a helper that can be used to export your releases.`,
		Example: heredoc.Doc(`
			$ rex generate changelog
			$ rex auth login
		`),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if f.Debug {
				logrus.SetLevel(logrus.DebugLevel)
			}
		},
	}

	formattedVersion := version.Format(appVersion, buildDate)

	cs := f.Colours()
	helpHelper := func(command *cobra.Command, args []string) {
		rootHelpFunc(cs, command, args)
	}

	rootCmd.SetHelpFunc(helpHelper)
	rootCmd.SetFlagErrorFunc(rootFlagErrorFunc)

	rootCmd.AddCommand(version.NewCmdVersion(appVersion, buildDate))

	rootCmd.SetVersionTemplate(formattedVersion)
	rootCmd.Version = formattedVersion
	rootCmd.Flags().Bool("version", false, "Show rex version")
	rootCmd.PersistentFlags().BoolVarP(&f.Debug, "debug", "d", false, "Debug output")

	rootCmd.SuggestionsMinimumDistance = 1

	return rootCmd
}
