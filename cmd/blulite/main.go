package main

import (
	"github.com/spf13/cobra"

	"github.com/PI-Victor/blu-lite/pkg/cmd"
)

func main() {
	RootCmd.AddCommand(cmd.StartCommand)
	RootCmd.Execute()
}

// RootCmd prints instructions about how to start the API Server.
var RootCmd = &cobra.Command{
	Use:   "blulite",
	Short: "Blu Lite API Server",
	Example: `
Specify the application working directory with bluelite start
--workdir=/path/to/dir. The application will try to find and validate the
config.yml file in the specified working directory. If validation fails, it
will exist with an error. If no config file is found, it will create one with
defaults.

You can start a new instance of the API Server by specyfing a config file
with bluelite start --config-file=/path/to/config/file.
If you don't, the application will create default settings and store them in
its working directory.
`,
	Run: func(command *cobra.Command, args []string) {
		command.Help()
	},
}
