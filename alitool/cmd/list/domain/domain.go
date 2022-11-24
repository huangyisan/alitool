/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package domain

import (
	"alitool/internal/ali/domain"
	"github.com/spf13/cobra"
)

func getDomainResource() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		domain.DoGetDomainResource()
	}
}

// domainCmd represents the domain command
var DomainCmd = &cobra.Command{
	Use:                   "domain",
	Short:                 "list all domains",
	DisableFlagsInUseLine: true,
	Example: `  # List all domains
  alitool list domain`,
	Run: getDomainResource(),
}

func init() {
	//cmd.rootCmd.AddCommand(domainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}