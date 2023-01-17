package sslcert

import (
	"alitool/internal/alego"
	"github.com/spf13/cobra"
)

var (
	detail bool
)

func sslCertAction() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if detail {
			alego.ListCertificatesDetail()
			return
		}
		alego.ListCertificates()
	}
}

// SslcertCmd represents the domain command
var SslcertCmd = &cobra.Command{
	Use:                   "sslcert",
	Short:                 "list all certificates",
	DisableFlagsInUseLine: true,
	Example: `  # List all domains
  alitool list sslcert`,
	Run: sslCertAction(),
}

func init() {
	SslcertCmd.Flags().BoolVarP(&detail, "detail", "d", false, "show certificates detail")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
