package commands

import (
	"github.com/spf13/cobra"
)

var containerCmd = &cobra.Command{
	Use:   "container",
	Short: "Container-related actions",
	Run: func(ccmd *cobra.Command, args []string) {
		ccmd.HelpFunc()(ccmd, args)
	},
}

func init() {
	containerCmd.AddCommand(containerGroupCmd)
	containerCmd.AddCommand(containerServiceCmd)
}
