package dns

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/common"
	. "alitool/internal/pkg/mylog"
)

// listDnsByAccount list dns by ali account
func (d *DnsClient) listDnsByAccount() recordDnsDomains {
	return d.getAllDnsDomains()
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

// ListDnsByAccount list dns by ali account
func ListDnsByAccount(i IDNSClient) {
	if account.IsExistAccount(i.getAccountName()) {
		hasRecordDomains := i.listDnsByAccount()
		LoggerNoT.Printf("%s has dns record:\n", i.getAccountName())
		for record, _ := range hasRecordDomains {
			LoggerNoT.Println(record)
		}
		return
	}
	LoggerNoT.Printf("%s is right?\n", i.getAccountName())
}

// IsDnsInAccount judge dns in account
func IsDnsInAccount(i IDNSClient, domainName string) {
	if account.IsExistAccount(i.getAccountName()) && common.IsValidDomain(domainName) {
		if ok := i.isDnsInAccount(domainName); ok {
			LoggerNoT.Printf("%s exist in %s\n", common.DomainSuffix(domainName), i.getAccountName())
			return
		}
		LoggerNoT.Printf("%s not exist in %s\n", domainName, i.getAccountName())
		return
	}
	LoggerNoT.Printf("invalid account: %s or domain: %s \n", i.getAccountName(), domainName)

}

// FindDnsInAccount find dns in which account
func FindDnsInAccount(domainName string) {
	initAllDnsClients()
	IDnsClients := getDnsClients()
	for _, v := range IDnsClients {
		if v.isDnsInAccount(domainName) {
			LoggerNoT.Printf("domain name %s in %s account\n", common.DomainSuffix(domainName), v.getAccountName())
			return
		}
	}
	LoggerNoT.Printf("domain name %s not in any accounts\n", common.DomainSuffix(domainName))
}
