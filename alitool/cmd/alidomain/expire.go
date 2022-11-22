/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package alidomain

import (
	domain "alitool/internal/domain"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	domainName string
)

// checkCmd represents the check command
var expire = &cobra.Command{
	Use:   "expire",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		regionId := viper.GetString("regionId")
		accessKeyId := viper.GetString("accessKeyId")
		accessKeySecret := viper.GetString("accessKeySecret")
		dc := domain.NewDomainClient(regionId, accessKeyId, accessKeySecret)
		dc.GetDomainExpireCurrDiff(domainName)
	},
}

func init() {
	DomainCmd.AddCommand(expire)
	expire.Flags().StringVarP(&domainName, "domainName", "d", "", "The domain to check")

	expire.Flags().String("all", "A", "check all domains")

	if err := expire.MarkFlagRequired("domainName"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
