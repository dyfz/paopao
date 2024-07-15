package serve

import (
	"github.com/spf13/cobra"
	"paopao/cmd"
)

func init() {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "start paopao-ce server",
		Long:  "start paopao-ce server",
		Run:   serveRun,
	}

	cmd.Register(serveCmd)
}

func serveRun(_ *cobra.Command, _ []string) {

}
