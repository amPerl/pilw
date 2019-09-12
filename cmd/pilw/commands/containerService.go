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

var containerServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Container service related actions",
	Run: func(ccmd *cobra.Command, args []string) {
		ccmd.HelpFunc()(ccmd, args)
	},
}

var containerServiceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List container services",
	Run:   containerServiceList,
}

var containerServiceRestartCmd = &cobra.Command{
	Use:   "restart [suuid]",
	Short: "Restart a container by its SUUID",
	Args:  cobra.MinimumNArgs(1),
	Run:   containerServiceRestart,
}

func init() {
	containerServiceListCmd.Flags().String("name", "", "Name of the service to filter by")
	containerServiceCmd.AddCommand(containerServiceListCmd)

	containerServiceCmd.AddCommand(containerServiceRestartCmd)
}

func printContainerServiceList(containerServiceList []api.ContainerService) {
	if viper.GetBool("quiet") {
		for _, containerService := range containerServiceList {
			fmt.Println(containerService.SUUID)
		}
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "SUUID\tNAME\tIMAGE\tCREATED AT")
	for _, containerService := range containerServiceList {
		fmt.Fprintln(w, fmt.Sprintf(
			"%s\t%s\t%s\t%s",
			containerService.SUUID,
			containerService.Name,
			containerService.Image,
			containerService.CreatedAt.Format(time.Stamp),
		))
	}
	w.Flush()
}

func containerServiceList(ccmd *cobra.Command, args []string) {
	apiKey := viper.GetString("key")

	containerServiceList, err := api.GetContainerServiceList(apiKey)
	if err != nil {
		log.Println("Failed to get container service list")
		log.Fatal(err)
	}

	nameFilter := ccmd.Flags().Lookup("name")
	if nameFilter.Changed {
		nameFilterString := nameFilter.Value.String()

		filteredContainerServiceList := make([]api.ContainerService, 0)
		for _, containerService := range containerServiceList {
			if containerService.Name == nameFilterString {
				filteredContainerServiceList = append(filteredContainerServiceList, containerService)
			}
		}

		containerServiceList = filteredContainerServiceList
	}

	printContainerServiceList(containerServiceList)
}

func containerServiceRestart(ccmd *cobra.Command, args []string) {
	apiKey := viper.GetString("key")
	suuid := args[0]

	err := api.StopContainerService(apiKey, suuid)
	if err != nil {
		log.Println("Failed to stop container service")
		log.Fatal(err)
	}

	err = api.StartContainerService(apiKey, suuid)
	if err != nil {
		log.Println("Failed to start container service")
		log.Fatal(err)
	}
}
