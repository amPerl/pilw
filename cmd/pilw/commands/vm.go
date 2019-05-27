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

var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "VM-related actions",
	Run: func(ccmd *cobra.Command, args []string) {
		ccmd.HelpFunc()(ccmd, args)
	},
}

var vmListCmd = &cobra.Command{
	Use:   "list",
	Short: "List VMs",
	Run:   vmList,
}

func init() {
	vmCmd.AddCommand(vmListCmd)
}

func printVMList(vmList []api.VM) {
	if viper.GetBool("quiet") {
		for _, vm := range vmList {
			fmt.Println(vm.ID)
		}
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tHOSTNAME\tDESCRIPTION\tPUBLIC IPV4\tMEMORY\tVCPU\tSTATUS\tCREATED AT")
	for _, vm := range vmList {
		fmt.Fprintln(w, fmt.Sprintf(
			"%d\t%s\t%s\t%s\t%s\t%d\t%d\t%s\t%s",
			vm.ID,
			vm.Name,
			vm.Description,
			vm.Hostname,
			vm.PublicIPv4,
			vm.Memory,
			vm.VCPU,
			vm.Status,
			vm.CreatedAt.Format(time.Stamp),
		))
	}
	w.Flush()
}

func vmList(ccmd *cobra.Command, args []string) {
	apiKey := viper.GetString("key")

	vmList, err := api.GetVMList(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	printVMList(vmList)
}
