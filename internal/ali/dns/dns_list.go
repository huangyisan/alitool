package dns

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/common"
)

// listDnsByAccount list dns by ali account
func listDnsByAccount(accountName string) RecordDomains {
	dnsClient := GetDnsClients()[accountName]
	return dnsClient.GetAllDomains()
	//if ok {
	//
	//}
	//return nil

}

// isDnsInAccount judege dns in account
func isDnsInAccount(accountName, domainName string) bool {
	_domainName := common.DomainSuffix(domainName)
	_, ok := listDnsByAccount(accountName)[_domainName]
	if ok {
		return true
	}
	return false
}

// findDnsInAccount reverse dns which ali account
func findDnsInAccount(domainName string) (accountName string) {
	_domainName := common.DomainSuffix(domainName)
	accountMap := account.GetAccountMap()
	for _accountName, _ := range accountMap {
		if _, ok := listDnsByAccount(_accountName)[_domainName]; ok {
			return _accountName
		}
	}
	return ""
}
