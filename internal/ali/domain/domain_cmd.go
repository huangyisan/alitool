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

// findDomainsInAccount reverse dns which ali account
func findDomainsInAccount(domainName string) (accountName string) {
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

// findExpireDomainsInAllAccounts will return expire domains in all accounts
func findExpireDomainsInAllAccounts(expireDay int) (expireDomainsInAllAccounts map[string]map[string]int) {
	expireDomainsInAllAccounts = make(map[string]map[string]int)
	accountMap := account.GetAccountMap()
	for _accountName, _ := range accountMap {
		expireDomainsInOneAccount := findExpireDomainsByAccount(_accountName, expireDay)
		if len(expireDomainsInOneAccount) > 0 {
			expireDomainsInAllAccounts[_accountName] = expireDomainsInOneAccount
		}
	}
	return expireDomainsInAllAccounts
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

// FindExpireDomainsByAccount will print expire domain in account
// alitool check  domain -a accountName -e 100
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

// FindExpireDomainsInAllAccounts will print all expire domains in every account
// alitool check  domain -A -e 100
func FindExpireDomainsInAllAccounts(expireDay int) {
	expireDomainsInAllAccounts := findExpireDomainsInAllAccounts(expireDay)
	if len(expireDomainsInAllAccounts) > 0 {
		for account, v := range expireDomainsInAllAccounts {
			if len(v) > 0 {
				fmt.Printf("account: %s", account)
				for domain, exDay := range v {
					fmt.Printf("domain: %s, expireDay: %s", domain, exDay)
				}
			}
		}
		return
	}
	fmt.Printf("no expire domain in %s days", expireDay)
}
