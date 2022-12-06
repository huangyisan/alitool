/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package domain

import (
	dm "alitool/internal/ali/domain"
	"github.com/spf13/cobra"
)

var (
	accountName string
	regionId    string
	domainName  string
	reverse     bool
)

func initDomainClient(accountName, regionId string) dm.IDomainClient {
	return dm.InitDomainClient(accountName, regionId)
}

func domainAction() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if accountName != "" && domainName == "" && reverse == false {
			domainClient := initDomainClient(accountName, regionId)
			dm.ListRegisteredDomainByAccount(domainClient)
			return
		}

		if accountName != "" && domainName != "" {
			domainClient := initDomainClient(accountName, regionId)
			dm.IsDomainInAccount(domainClient, domainName)
			return
		}

		if domainName != "" && reverse == true {
			dm.FindDomainInAccount(domainName)
			return
		}
		cmd.Help()

	}
}

// domainCmd represents the domain command
var DomainCmd = &cobra.Command{
	Use:                   "domain",
	Short:                 "list all domains",
	DisableFlagsInUseLine: true,
	Example: `  # List all domains
  alitool list domain`,
	Run: domainAction(),
}

func init() {

	DomainCmd.Flags().StringVarP(&domainName, "domain", "i", "", "specific domain name")
	DomainCmd.Flags().StringVarP(&accountName, "account", "a", "", "specific account name")
	DomainCmd.Flags().StringVarP(&regionId, "regionId", "z", "", "specific region id")
	DomainCmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "reverse the domain in account")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
