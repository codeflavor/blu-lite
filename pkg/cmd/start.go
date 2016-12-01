package cmd

import (
	"github.com/spf13/cobra"

	"github.com/codeflavor/servops/pkg/cmd/start"
)

var (
	confFile string
	workDir  string
	logLevel string
)

// StartCmd launches a new instance of the app.
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the application",
	Run: func(command *cobra.Command, args []string) {
		start.ValidateUserConfig(confFile, workDir, logLevel)
	},
}

func init() {
	StartCmd.PersistentFlags().StringVar(&workDir, "dir", "", "Specify the application working directory")
	StartCmd.PersistentFlags().StringVar(&confFile, "conf", "", "Specify a valid configuration file")
	StartCmd.PersistentFlags().StringVar(&logLevel, "loglevel", "0", "Specify an output loglevel")
}
