package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	confFile string
	workDir  string
)

// StartCommand launches a new instance of the app.
var StartCommand = &cobra.Command{
	Use:   "start",
	Short: "Start the application",
	Run: func(command *cobra.Command, args []string) {
		fmt.Println("Started")
	},
}

func init() {
	StartCommand.PersistentFlags().StringVar(&workDir, "workdir", "", "Specify the application working directory")
	StartCommand.PersistentFlags().StringVar(&confFile, "config", "", "Specify a valid configuration file")
}
