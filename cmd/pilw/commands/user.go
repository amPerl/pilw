package commands

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/amPerl/pilw/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User-related actions",
	Run: func(ccmd *cobra.Command, args []string) {
		ccmd.HelpFunc()(ccmd, args)
	},
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display authenticated user info",
	Run:   info,
}

func init() {
	userCmd.AddCommand(infoCmd)
}

func info(ccmd *cobra.Command, args []string) {
	apiKey := viper.GetString("key")

	userInfo, err := api.GetUserInfo(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	if viper.GetBool("quiet") {
		fmt.Println(userInfo.ID)
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "ID\tCOOKIE_ID\tNAME\tLAST_ACTIVITY")
	fmt.Fprintln(w, fmt.Sprintf(
		"%d\t%s\t%s\t%s",
		userInfo.ID,
		userInfo.CookieID,
		userInfo.Name,
		userInfo.LastActivity.String(),
	))
	w.Flush()
}
