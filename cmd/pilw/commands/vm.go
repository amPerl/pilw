package commands

import (
	"fmt"
	"log"
	"net/url"
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

var vmUpdateCmd = &cobra.Command{
	Use:   "update [uuid]",
	Short: "Update one or more fields on a VM",
	Args:  cobra.MinimumNArgs(1),
	Run:   vmUpdate,
}

func init() {
	vmCmd.AddCommand(vmListCmd)

	vmUpdateCmd.Flags().String("name", "", "Name of the VM")
	vmUpdateCmd.Flags().Int("ram", 0, "RAM in megabytes (has to be set with vCPU)")
	vmUpdateCmd.Flags().Int("vcpu", 0, "vCPU in cores (has to be set with RAM)")
	vmCmd.AddCommand(vmUpdateCmd)
}

func printVMList(vmList []api.VM) {
	if viper.GetBool("quiet") {
		for _, vm := range vmList {
			fmt.Println(vm.UUID)
		}
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "UUID\tNAME\tDESCRIPTION\tHOSTNAME\tOS\tPUBLIC IPV4\tMEMORY\tVCPU\tSTORAGE\tSTATUS\tCREATED AT")
	for _, vm := range vmList {
		fmt.Fprintln(w, fmt.Sprintf(
			"%s\t%s\t%s\t%s\t%s\t%s\t%d\t%d\t%d\t%s\t%s",
			vm.UUID,
			vm.Name,
			vm.Description,
			vm.Hostname,
			fmt.Sprintf("%s %s", vm.OSName, vm.OSVersion),
			vm.PublicIPv4,
			vm.Memory,
			vm.VCPU,
			len(vm.Storage),
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

func vmUpdate(ccmd *cobra.Command, args []string) {
	apiKey := viper.GetString("key")
	uuid := args[0]
	changedFields := url.Values{}

	nameFlag := ccmd.Flags().Lookup("name")
	if nameFlag.Changed {
		changedFields.Add("name", nameFlag.Value.String())
	}

	ramFlag := ccmd.Flags().Lookup("ram")
	if ramFlag.Changed {
		changedFields.Add("ram", ramFlag.Value.String())
	}

	vCPUFlag := ccmd.Flags().Lookup("vcpu")
	if vCPUFlag.Changed {
		changedFields.Add("vcpu", vCPUFlag.Value.String())
	}

	err := api.UpdateVM(apiKey, uuid, changedFields)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(uuid)
}
