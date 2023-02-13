/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package list

import (
	"alitool/alitool/cmd/list/account"
	"alitool/alitool/cmd/list/dcdn"
	"alitool/alitool/cmd/list/dns"
	"alitool/alitool/cmd/list/domain"
	"alitool/alitool/cmd/list/sslcert"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var ListCmd = &cobra.Command{
	Use:                   "list",
	Short:                 "list resource",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	ListCmd.AddCommand(domain.DomainCmd)
	ListCmd.AddCommand(account.AccountCmd)
	ListCmd.AddCommand(account.RegionCmd)
	ListCmd.AddCommand(dns.DnsCmd)
	ListCmd.AddCommand(sslcert.SslcertCmd)
	ListCmd.AddCommand(dcdn.DcdnCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
