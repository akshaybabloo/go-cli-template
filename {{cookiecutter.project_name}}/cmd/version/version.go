package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdVersion(version, buildDate string) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "version",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(Format(version, buildDate))
		},
	}

	return cmd
}

func Format(version, buildDate string) string {
	return fmt.Sprintf("%s %s\n", version, buildDate)
}
