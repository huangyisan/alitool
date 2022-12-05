/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package domain

import (
	"alitool/internal/ali/domain"
	"github.com/spf13/cobra"
)

var (
	domainName     string
	accountName    string
	checkAllDomain bool
	expireDay      int
)

func domainAction() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if checkAllDomain {
			domain.FindExpireDomainsInAllAccounts(expireDay)
			return
		}

		if accountName != "" && domainName == "" {
			domain.FindExpireDomainsByAccount(accountName, expireDay)
			return
		}

		if domainName != "" && accountName == "" {
			domain.FindExpireDomainRefAccount(domainName)
			return
		}
		DomainCmd.Help()
	}
}

// domainCmd represents the domain command
var DomainCmd = &cobra.Command{
	Use:   "domain",
	Short: "check domain expire day",
	Example: `  # check all domains expire day
  alitool check domain -A -e 100

  # check specific domain within end expire day
  alitool check domain -d baidu.com 
`,
	Run: domainAction(),
}

func init() {
	DomainCmd.Flags().StringVarP(&accountName, "account", "a", "", "specific account to check")
	DomainCmd.Flags().StringVarP(&domainName, "domain", "d", "", "specific domain to check")
	DomainCmd.Flags().IntVarP(&expireDay, "end-expire-day", "e", 100, "specific end expire day")
	DomainCmd.Flags().BoolVarP(&checkAllDomain, "all-domains", "A", false, "check all domains")

	// only one command execute at one time
	DomainCmd.MarkFlagsMutuallyExclusive("all-domains", "domain")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
