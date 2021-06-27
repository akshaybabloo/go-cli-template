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
		Use:   "{{cookiecutter.project_name}} [OPTIONS] [COMMANDS]",
		Short: "{{cookiecutter.short_description}}",
		Long:  `{{cookiecutter.project_name}} {{cookiecutter.long_description}}.`,
		Example: heredoc.Doc(`
			$ {{cookiecutter.project_name}} --version
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
	rootCmd.Flags().Bool("version", false, "Show {{cookiecutter.project_name}} version")
	rootCmd.PersistentFlags().BoolVarP(&f.Debug, "debug", "d", false, "Debug output")

	rootCmd.SuggestionsMinimumDistance = 1

	return rootCmd
}
