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

// StartCommand launches a new instance of the app.
var StartCommand = &cobra.Command{
	Use:   "start",
	Short: "Start the application",
	Run: func(command *cobra.Command, args []string) {
		start.ValidateUserConfig(confFile, workDir, logLevel)
	},
}

func init() {
	StartCommand.PersistentFlags().StringVar(&workDir, "dir", "", "Specify the application working directory")
	StartCommand.PersistentFlags().StringVar(&confFile, "conf", "", "Specify a valid configuration file")
	StartCommand.PersistentFlags().StringVar(&logLevel, "loglevel", "0", "Specify an output loglevel")
}
