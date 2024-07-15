package cmd

import (
	"github.com/spf13/cobra"
	"paopao/pkg/utils"
	"paopao/version"
)

func init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "show version information",
		Long:  "show version information",
		Run:   versionRun,
	}
	Register(versionCmd)
}

func versionRun(_ *cobra.Command, _ []string) {
	utils.PrintHelloBanner(version.VersionInfo())
}
