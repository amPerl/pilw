package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// PilwCmd is the entrypoint and root-level command for the CLI
	PilwCmd = &cobra.Command{
		Use:   "pilw",
		Short: "",
		Long:  ``,
		Run: func(ccmd *cobra.Command, args []string) {
			ccmd.HelpFunc()(ccmd, args)
		},
	}
)

func init() {
	PilwCmd.PersistentFlags().BoolP("quiet", "q", false, "Restrict output to bare minimum (usually IDs)")

	viper.BindPFlag("quiet", PilwCmd.PersistentFlags().Lookup("quiet"))

	PilwCmd.AddCommand(userCmd)
	PilwCmd.AddCommand(tokenCmd)
}
