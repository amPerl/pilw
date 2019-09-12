package commands

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/amPerl/pilw/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var containerGroupCmd = &cobra.Command{
	Use:   "group",
	Short: "Container group related actions",
	Run: func(ccmd *cobra.Command, args []string) {
		ccmd.HelpFunc()(ccmd, args)
	},
}

var containerGroupListCmd = &cobra.Command{
	Use:   "list",
	Short: "List container groups",
	Run:   containerGroupList,
}

func init() {
	containerGroupCmd.AddCommand(containerGroupListCmd)
}

func printContainerGroupList(containerGroupList []api.ContainerGroup) {
	if viper.GetBool("quiet") {
		for _, containerGroup := range containerGroupList {
			fmt.Println(containerGroup.SUUID)
		}
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "SUUID\tNAME\tSERVICES\tCREATED AT")
	for _, containerGroup := range containerGroupList {
		fmt.Fprintln(w, fmt.Sprintf(
			"%s\t%s\t%d\t%s",
			containerGroup.SUUID,
			containerGroup.Name,
			len(containerGroup.Services),
			containerGroup.CreatedAt.Format(time.Stamp),
		))
	}
	w.Flush()
}

func containerGroupList(ccmd *cobra.Command, args []string) {
	apiKey := viper.GetString("key")

	containerGroupList, err := api.GetContainerGroupList(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	printContainerGroupList(containerGroupList)
}
