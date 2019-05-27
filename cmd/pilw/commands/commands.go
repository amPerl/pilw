package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	PilwCmd = &cobra.Command{
		Use:   "",
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
}
