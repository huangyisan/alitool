/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package sslcert

import (
	"alitool/internal/alego"
	"github.com/go-acme/lego/v4/lego"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	domainNames []string
	platform    string
	testFlag    bool
)

func sslCertAction() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		client := &lego.Client{}
		acmeEmail := viper.GetString("acme.account.name")
		acmeApi := ""
		if testFlag {
			acmeApi = viper.GetString("acme.api_url.staging")
			client = alego.NewAcmeClient(acmeEmail, acmeApi)
		} else {
			acmeApi = viper.GetString("acme.api_url.prod")
			client = alego.NewAcmeClient(acmeEmail, acmeApi)
		}

		switch platform {
		case "cloudflare":
			authEmail := viper.GetString("acme.ssl_platform.cloudflare.authEmail")
			authToken := viper.GetString("acme.ssl_platform.cloudflare.authToken")
			alego.CloudFlareVerification(client, authEmail, authToken, domainNames...)
		default:
			cmd.Help()
		}
	}
}

// SslcertCmd represents the sslcert command
var SslcertCmd = &cobra.Command{
	Use:   "sslcert",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: sslCertAction(),
}

func init() {
	SslcertCmd.Flags().StringSliceVarP(&domainNames, "domain", "d", []string{}, "specify ssl domain")
	SslcertCmd.Flags().StringVarP(&platform, "platform", "p", "cloudflare", "specify verification cloud platform")
	SslcertCmd.Flags().BoolVarP(&testFlag, "test", "t", false, "specific test api")

	//cmd.rootCmd.AddCommand(sslcertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sslcertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sslcertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
