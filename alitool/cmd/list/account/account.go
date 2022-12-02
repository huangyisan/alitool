/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package account

import (
	"alitool/internal/ali/account"
	"github.com/spf13/cobra"
)

func getAccountResource() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		account.ListAccount()
	}
}

func listRegionResource() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		account.ListRegion()
	}
}

// accountCmd represents the account command
var AccountCmd = &cobra.Command{
	Use:                   "account",
	Short:                 "List all ali accounts",
	DisableFlagsInUseLine: true,
	Example: `  # List all domains
    alitool list account`,
	Run: getAccountResource(),
}

var RegionCmd = &cobra.Command{
	Use:                   "region",
	Short:                 "List all region",
	DisableFlagsInUseLine: true,
	Example: `  # List all region
    alitool list region`,
	Run: listRegionResource(),
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
