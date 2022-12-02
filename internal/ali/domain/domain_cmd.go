package domain

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/common"
	"fmt"
)

// listRegisteredDomainByAccount list registered domain by ali account
func listRegisteredDomainByAccount(accountName string) recordRegisterDomains {
	domainClient := GetDomainClients()[accountName]
	return domainClient.getAllRegisteredDomains()
}

// isDomainInAccount judge registered domain in account
func isDomainInAccount(accountName, domainName string) bool {
	_domainName := common.DomainSuffix(domainName)
	_, ok := listRegisteredDomainByAccount(accountName)[_domainName]
	if ok {
		return true
	}
	return false
}

// findDomainInAccount reverse dns which ali account
func findDomainInAccount(domainName string) (accountName string) {
	_domainName := common.DomainSuffix(domainName)
	accountMap := account.GetAccountMap()
	for _accountName, _ := range accountMap {
		if _, ok := listRegisteredDomainByAccount(_accountName)[_domainName]; ok {
			return _accountName
		}
	}
	return ""
}

// findExpireDomainsByAccount will return expire domain in specific day and account
// expireDomains map[domainName]willExpireDay
func findExpireDomainsByAccount(accountName string, expireDay int) (expireDomains map[string]int) {
	domainClient := GetDomainClients()[accountName]
	return domainClient.getExpireDomains(expireDay)
}

func ListRegisteredDomainByAccount(accountName string) {
	recordRegisterDomains := listRegisteredDomainByAccount(accountName)
	if len(recordRegisterDomains) > 0 {
		fmt.Printf("account %s exist registed domain\n", accountName)
		for d, _ := range recordRegisterDomains {
			fmt.Printf("domain: %s\n", d)
		}
	}
	fmt.Printf("Total count: %d\n", len(recordRegisterDomains))
}

func FindExpireDomainsByAccount(accountName string, expireDay int) {
	expireDomains := findExpireDomainsByAccount(accountName, expireDay)
	if len(expireDomains) > 0 {
		for d, e := range expireDomains {
			fmt.Printf("account %s domain %s will expire in %d\n", accountName, d, e)
		}
		return
	}
	fmt.Printf("account %s no expire domain in %d days\n", accountName, expireDay)
}
