package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = cobra.Command{
		Use:   "paopao",
		Short: `an artistic "twitter like" community`,
		Long:  `an artistic "twitter like" community`,
	}
)

func Register(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func Execute() {
	rootCmd.Execute()
}
