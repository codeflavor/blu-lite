package main

import (
	"github.com/spf13/cobra"

	"github.com/codeflavor/servops/pkg/cmd"
)

// RootCmd prints instructions about how to start the API Server.
var RootCmd = &cobra.Command{
	Use:   "servops",
	Short: "servops API Server",
	Example: `
Specify the application working directory with servops start
--workdir=/path/to/dir. The application will try to find and validate the
config.yml file in the specified working directory. If validation fails, it
will exist with an error. If no config file is found, it will generate one with
defaults.
`,
	Run: func(command *cobra.Command, args []string) {
		command.Help()
	},
}

func main() {
	RootCmd.AddCommand(cmd.StartCommand)
	RootCmd.Execute()
}
