/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package ssl

import (
	"alitool/internal/ali/ssl"
	"github.com/spf13/cobra"
)

var (
	accountName string
	regionId    string
	allCerts    bool
	expireDay   int
)

func initSSLClient(accountName, regionId string) ssl.ISSLClient {
	return ssl.InitSSLClient(accountName, regionId)
}

func sslAction() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if allCerts {
			ssl.GetExpireCertInAllAccounts(expireDay)
			ssl.GetExpireUploadCertInAllAccounts(expireDay)
			return
		}
		if accountName != "" {
			sslClient := initSSLClient(accountName, regionId)
			ssl.GetExpireCertByAccount(sslClient, expireDay)
			ssl.GetExpireUploadCertByAccount(sslClient, expireDay)
			return
		}
		cmd.Help()
	}
}

// SSLCmd represents the ssl command
var SSLCmd = &cobra.Command{
	Use:   "ssl",
	Short: "check certs expire day",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: sslAction(),
}

func init() {
	SSLCmd.Flags().StringVarP(&accountName, "account", "a", "", "specific account to check")
	SSLCmd.Flags().StringVarP(&regionId, "regionId", "z", "cn-shanghai", "specific account region id")
	SSLCmd.Flags().IntVarP(&expireDay, "end-expire-day", "e", 30, "specific end expire day")
	SSLCmd.Flags().BoolVarP(&allCerts, "all-certs", "A", false, "check all certs")

	// only one command execute at one time
	//SSLCmd.MarkFlagsMutuallyExclusive("all-certs", "account")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sslCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sslCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
