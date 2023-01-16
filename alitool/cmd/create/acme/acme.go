/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package acme

import (
	"alitool/internal/alego"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	accountName string
	testFlag    bool
	prodApi     string
	stagingApi  string
)

func acmeAction() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		logrus.Infof("accountName: %s\n", accountName)
		if accountName != "" && testFlag {
			alego.RegistrationAcmeAccount(accountName, stagingApi)
			return
		}
		if accountName != "" && !testFlag {
			alego.RegistrationAcmeAccount(accountName, prodApi)
			return
		}
		cmd.Help()
	}
}

// AcmeCmd represents the acme command
var AcmeCmd = &cobra.Command{
	Use:   "acme",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: acmeAction(),
}

func InitValues() {
	// Get value from ~/.alitool.yaml
	prodApi = viper.GetString("acme.api_url.prod")
	stagingApi = viper.GetString("acme.api_url.staging")
	accountName = viper.GetString("acme.account.name")
}

func init() {
	//AcmeCmd.Flags().StringVarP(&accountName, "user", "u", "", "specific account name")
	AcmeCmd.Flags().BoolVarP(&testFlag, "test", "t", false, "specific test api")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// acmeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// acmeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
