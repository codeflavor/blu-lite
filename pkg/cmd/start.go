package cmd

import (
	"github.com/spf13/cobra"

	c "github.com/codeflavor/servops/pkg/config"
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
		c.LoadConfig(workDir, logLevel)
	},
}

func init() {
	StartCmd.PersistentFlags().StringVar(&workDir, "dir", "", "Specify the application working directory")
	StartCmd.PersistentFlags().StringVar(&logLevel, "loglevel", "0", "Specify an output loglevel")
}
