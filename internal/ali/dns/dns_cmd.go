package dns

import (
	"alitool/internal/pkg/common"
	"fmt"
)

// ListDnsByAccount list dns by ali account
func ListDnsByAccount(accountName string) {
	if common.IsExistAccount(accountName) {
		hasRecordDomains := listDnsByAccount(accountName)
		fmt.Printf("%s has dns record:\n", accountName)
		for record, _ := range hasRecordDomains {
			fmt.Println(record)
		}
		return
	}
	fmt.Printf("%s is right?\n", accountName)
}

// IsDnsInAccount judge dns in account
func IsDnsInAccount(accountName, domainName string) {
	if common.IsExistAccount(accountName) && common.IsValidDomain(domainName) {
		if ok := isDnsInAccount(accountName, domainName); ok {
			fmt.Printf("%s exist in %s\n", common.DomainSuffix(domainName), accountName)
			return
		}
		fmt.Printf("%s not exist in %s\n", domainName, accountName)
		return
	}
	fmt.Printf("invalid account: %s or domain: %s \n", accountName, domainName)

}

// FindDnsInAccount find dns in which account
func FindDnsInAccount(domainName string) {
	accountName := findDnsInAccount(domainName)
	if accountName != "" {
		fmt.Printf("domain name %s in %s account\n", common.DomainSuffix(domainName), accountName)
		return
	}
	fmt.Printf("domain name %s not in any accounts\n", common.DomainSuffix(domainName))
}
