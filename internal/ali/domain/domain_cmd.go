package domain

import (
	"alitool/internal/pkg/common"
	"fmt"
)

// listRegisteredDomainByAccount list registered domain by ali account
func (d *DomainClient) listRegisteredDomainByAccount() recordRegisterDomains {
	return d.getAllRegisteredDomains()
}

// isDomainInAccount judge registered domain in account
func (d *DomainClient) isDomainInAccount(domainName string) bool {
	_domainName := common.DomainSuffix(domainName)
	_, ok := d.listRegisteredDomainByAccount()[_domainName]
	if ok {
		return true
	}
	return false
}

// IsDomainInAccount print domain whether in account
func IsDomainInAccount(i IDomainClient, domainName string) {
	if common.IsExistAccount(i.getAccountName()) && common.IsValidDomain(domainName) {
		if i.isDomainInAccount(domainName) {
			fmt.Printf("%s exist in %s", domainName, i.getAccountName())
			return
		}
	}
	fmt.Printf("%s not exist in %s", domainName, i.getAccountName())
}

func FindDomainInAccount(domainName string) {
	initAllDomainClient()
	IDomainClients := getDomainClients()
	for _, v := range IDomainClients {
		if v.isDomainInAccount(domainName) {
			fmt.Printf("domainName: %s exist in %s", domainName, v.getAccountName())
			return
		}
	}
	fmt.Printf("domainName: %s not exist in any account", domainName)
}

// findExpireDomainsByAccount will return expire domain in specific day and account
// expireDomains map[domainName]willExpireDay
func (d *DomainClient) findExpireDomainsByAccount(i IDomainClient, expireDay int) (expireDomains map[string]int) {
	return i.getExpireDomains(expireDay)
}

// findExpireDomainRefAccount will return expire domain and account
// e.g. alitool check  domain -d baidu.com
func (d *DomainClient) findExpireDomainRefAccount(domainName string) (accountName string, expireDay int) {
	if d.isDomainInAccount(domainName) {
		return d.getAccountName(), d.getDomainExpireCurrDiff(domainName)
	}
	return d.getAccountName(), -1
}

func FindExpireDomainRefAccount(i IDomainClient, domainName string) {
	accountName, expireDay := i.findExpireDomainRefAccount(domainName)
	if accountName != "" {
		fmt.Printf("domain %s found in %s account, expire in %d days\n", domainName, accountName, expireDay)
		return
	}
	fmt.Printf("no %s found in all account", domainName)

}

func ListRegisteredDomainByAccount(i IDomainClient) {
	recordRegisterDomains := i.listRegisteredDomainByAccount()
	if len(recordRegisterDomains) > 0 {
		fmt.Printf("account %s exist registed domain\n", i.getAccountName())
		for d, _ := range recordRegisterDomains {
			fmt.Printf("domain: %s\n", d)
		}
	}
	fmt.Printf("Total count: %d\n", len(recordRegisterDomains))
}

// FindExpireDomainsByAccount will print expire domain in account
// alitool check  domain -a AccountName -e 100
func FindExpireDomainsByAccount(i IDomainClient, expireDay int) {
	expireDomains := i.findExpireDomainsByAccount(i, expireDay)
	if len(expireDomains) > 0 {
		for d, e := range expireDomains {
			fmt.Printf("account %s domain %s will expire in %d\n", i.getAccountName(), d, e)
		}
		return
	}
	fmt.Printf("account %s no expire domain in %d days\n", i.getAccountName(), expireDay)
}

// FindExpireDomainsInAllAccounts will print all expire domains in every account
// alitool check  domain -A -e 100
func FindExpireDomainsInAllAccounts(expireDay int) {
	initAllDomainClient()
	IDomainClients := getDomainClients()
	for _, v := range IDomainClients {
		FindExpireDomainsByAccount(v, expireDay)
	}
}
