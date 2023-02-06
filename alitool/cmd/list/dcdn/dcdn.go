/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package dcdn

import (
	"alitool/internal/ali/dcdn"
	"github.com/spf13/cobra"
)

var (
	accountName string
	allDomains  bool
	regionId    string
)

func createDcdnClient(accountName, regionId string) dcdn.IDcdnClient {
	return dcdn.InitDcdnClient(accountName, regionId)
}

func dcdnAction() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if allDomains && accountName == "" {
			dcdn.ListAllHttpsDomainsInAllAccount()
			return
		}
		if !allDomains && accountName != "" {
			idcdn := createDcdnClient(accountName, regionId)
			dcdn.ListAllHttpsDomainsByAccount(idcdn)
			return
		}
		cmd.Help()
	}
}

// dcdnCmd represents the dcdn command
var DcdnCmd = &cobra.Command{
	Use:   "dcdn",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: dcdnAction(),
}

func init() {
	// Here you will define your flags and configuration settings.
	DcdnCmd.Flags().StringVarP(&accountName, "account", "a", "", "specified account name")
	DcdnCmd.Flags().BoolVarP(&allDomains, "all-domains", "A", false, "check all domains")
	DcdnCmd.Flags().StringVarP(&regionId, "region", "z", "cn-shanghai", "specific account region")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dcdnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dcdnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
