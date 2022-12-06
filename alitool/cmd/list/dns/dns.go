/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package dns

import (
	_dns "alitool/internal/ali/dns"
	"github.com/spf13/cobra"
)

var (
	domainName  string
	accountName string
	regionId    string
	reverse     bool
)

func createDNSClient(accountName, regionId string) _dns.IDNSClient {
	return _dns.InitDnsClient(accountName, regionId)
}

func dnsAction() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {

		//exist reverse flag
		if reverse && domainName != "" {
			_dns.FindDnsInAccount(domainName)
			return
		}

		if accountName != "" && domainName != "" {
			iDns := createDNSClient(accountName, regionId)
			_dns.IsDnsInAccount(iDns, domainName)
			return
		}

		if accountName != "" && domainName == "" {
			iDns := createDNSClient(accountName, regionId)
			_dns.ListDnsByAccount(iDns)
			return
		}
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
	DnsCmd.Flags().StringVarP(&regionId, "region", "z", "cn-shanghai", "specific account region")
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
