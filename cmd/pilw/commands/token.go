package commands

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/amPerl/pilw/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Token-related actions",
	Run: func(ccmd *cobra.Command, args []string) {
		ccmd.HelpFunc()(ccmd, args)
	},
}

var tokenListCmd = &cobra.Command{
	Use:   "list",
	Short: "List user API tokens",
	Run:   tokenList,
}

var tokenCreateCmd = &cobra.Command{
	Use:   "create [description] [restricted=false] [billing_account_id=0]",
	Short: "Create new token and register it at API Gateway",
	Args:  cobra.MinimumNArgs(1),
	Run:   tokenCreate,
}

var tokenDeleteCmd = &cobra.Command{
	Use:   "delete [token_id]",
	Short: "Delete a token by its ID",
	Args:  cobra.MinimumNArgs(1),
	Run:   tokenDelete,
}

func init() {
	tokenCmd.AddCommand(tokenListCmd)
	tokenCmd.AddCommand(tokenCreateCmd)
	tokenCmd.AddCommand(tokenDeleteCmd)
}

func printTokenList(tokenList []api.Token) {
	if viper.GetBool("quiet") {
		for _, token := range tokenList {
			fmt.Println(token.ID)
		}
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "ID\tDESCRIPTION\tRESTRICTED\tCREATED_AT\tBILLING_ACCOUNT_ID")
	for _, token := range tokenList {
		fmt.Fprintln(w, fmt.Sprintf(
			"%d\t%s\t%v\t%s\t%d",
			token.ID,
			token.Description,
			token.Restricted,
			token.CreatedAt.Format(time.Stamp),
			token.BillingAccountID,
		))
	}
	w.Flush()
}

func tokenList(ccmd *cobra.Command, args []string) {
	apiKey := viper.GetString("key")

	tokenList, err := api.GetTokenList(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	printTokenList(tokenList)
}

func tokenCreate(ccmd *cobra.Command, args []string) {
	apiKey := viper.GetString("key")

	var err error
	description := args[0]
	restricted := false
	billingAccountID := int64(0)

	if len(args) > 1 {
		restricted, err = strconv.ParseBool(args[1])
		if err != nil {
			log.Fatalf("\"%s\" is not a valid boolean value", args[1])
		}
	}

	if len(args) > 2 {
		billingAccountID, err = strconv.ParseInt(args[2], 10, 32)
		if err != nil {
			log.Fatalf("\"%s\" is not a valid integer value", args[2])
		}
	}

	tokenList, err := api.CreateToken(apiKey, description, restricted, int(billingAccountID))
	if err != nil {
		log.Fatal(err)
	}

	printTokenList(tokenList)
}

func tokenDelete(ccmd *cobra.Command, args []string) {
	apiKey := viper.GetString("key")

	tokenID, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		log.Fatalf("\"%s\" is not a valid integer value", args[0])
	}

	err = api.DeleteToken(apiKey, int(tokenID))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tokenID)
}
