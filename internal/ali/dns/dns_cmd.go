package dns

import (
	"alitool/internal/pkg/common"
	"fmt"
)

// listDnsByAccount list dns by ali account
func (d *DnsClient) listDnsByAccount() recordDnsDomains {
	return d.getAllDnsDomains()
	//dnsClient := GetDnsClients()[accountName]
	//return dnsClient.getAllDnsDomains()

}

// isDnsInAccount judge dns in account
func (d *DnsClient) isDnsInAccount(domainName string) bool {
	_domainName := common.DomainSuffix(domainName)
	_, ok := d.listDnsByAccount()[_domainName]
	if ok {
		return true
	}
	return false
}

// findDnsInAccount reverse dns which ali account
//func findDnsInAccount(domainName string) (accountName string) {
//	_domainName := common.DomainSuffix(domainName)
//	accountMap := account.GetAccountMap()
//	for _accountName, _ := range accountMap {
//		if _, ok := listDnsByAccount(_accountName)[_domainName]; ok {
//			return _accountName
//		}
//	}
//	return ""
//}

// ListDnsByAccount list dns by ali account

func ListDnsByAccount(i IDNSClient) {
	if common.IsExistAccount(i.getAccountName()) {
		hasRecordDomains := i.listDnsByAccount()
		fmt.Printf("%s has dns record:\n", i.getAccountName())
		for record, _ := range hasRecordDomains {
			fmt.Println(record)
		}
		return
	}
	fmt.Printf("%s is right?\n", i.getAccountName())
}

// IsDnsInAccount judge dns in account
func IsDnsInAccount(i IDNSClient, domainName string) {
	if common.IsExistAccount(i.getAccountName()) && common.IsValidDomain(domainName) {
		if ok := i.isDnsInAccount(domainName); ok {
			fmt.Printf("%s exist in %s\n", common.DomainSuffix(domainName), i.getAccountName())
			return
		}
		fmt.Printf("%s not exist in %s\n", domainName, i.getAccountName())
		return
	}
	fmt.Printf("invalid account: %s or domain: %s \n", i.getAccountName(), domainName)

}

// FindDnsInAccount find dns in which account
func FindDnsInAccount(domainName string) {
	initAllDnsClients()
	IDnsClients := getDnsClients()
	for _, v := range IDnsClients {
		if v.isDnsInAccount(domainName) {
			fmt.Printf("domain name %s in %s account\n", common.DomainSuffix(domainName), v.getAccountName())
			return
		}
	}

	fmt.Printf("domain name %s not in any accounts\n", common.DomainSuffix(domainName))
}
