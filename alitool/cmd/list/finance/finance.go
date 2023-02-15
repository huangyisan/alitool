/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package finance

import (
	fi "alitool/internal/ali/finance"
	"github.com/spf13/cobra"
)

var (
	accountName string
	regionId    string
	allFinances bool
)

func createFinanceClient(accountName, regionId string) fi.IFinanceClient {
	return fi.InitFinanceClient(accountName, regionId)
}

func financeAction() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if allFinances && accountName == "" {
			fi.ListLastMonthPaymentAmountInAllAccounts()
			return
		}
		if !allFinances && accountName != "" {
			account := createFinanceClient(accountName, regionId)
			fi.ListLastMonthPaymentAmountByAccount(account)
			return
		}
		cmd.Help()
	}
}

// financeCmd represents the finance command
var FinanceCmd = &cobra.Command{
	Use:   "finance",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: financeAction(),
}

func init() {
	//cmd.financeCmd.AddCommand(financeCmd)
	FinanceCmd.Flags().StringVarP(&accountName, "account", "a", "", "specific account name")
	FinanceCmd.Flags().StringVarP(&regionId, "region", "z", "cn-shanghai", "specific account region")
	FinanceCmd.Flags().BoolVarP(&allFinances, "all-domains", "A", false, "check all account finance")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// financeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// financeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
