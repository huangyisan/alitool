/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package dns

import (
	_dns "alitool/internal/ali/dns"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	domainName  string
	accountName string
	reverse     bool
)

func dnsAction() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		//exist reverse flag
		if reverse && domainName != "" {
			_dns.FindDnsInAccount(domainName)
			return
		}

		if accountName != "" && domainName != "" {
			_dns.IsDnsInAccount(accountName, domainName)
			return
		}

		if accountName != "" && domainName == "" {
			_dns.ListDnsByAccount(accountName)
			return
		}

		fmt.Println("tmp: exec no")
		fmt.Println(accountName)
		fmt.Println("tmp: exec no")
		cmd.Help()
	}
}

// dnsCmd represents the dns command
var DnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: dnsAction(),
}

func init() {
	// Here you will define your flags and configuration settings.
	DnsCmd.Flags().StringVarP(&domainName, "domain", "i", "", "specific domain name")
	DnsCmd.Flags().StringVarP(&accountName, "account", "a", "", "specific account name")
	DnsCmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "reverse the domain in account")

	// only one command execute at one time
	//DnsCmd.MarkFlagsMutuallyExclusive("account", "reverse")
	//DnsCmd.MarkFlagsRequiedTogether("domain", "reverse")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dnsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
