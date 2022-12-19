package domain

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/common"
	. "alitool/internal/pkg/mylog"
	"github.com/sirupsen/logrus"
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
	if account.IsExistAccount(i.getAccountName()) && common.IsValidDomain(domainName) {
		if i.isDomainInAccount(domainName) {
			LoggerNoT.Printf("%s exist in %s", domainName, i.getAccountName())
			return
		}
	}
	LoggerNoT.Printf("%s not exist in %s", domainName, i.getAccountName())
}

func FindDomainInAccount(domainName string) {
	initAllDomainClient()
	IDomainClients := getDomainClients()
	for _, v := range IDomainClients {
		if v.isDomainInAccount(domainName) {
			LoggerNoT.Printf("domainName: %s exist in %s", domainName, v.getAccountName())
			return
		}
	}
	LoggerNoT.Printf("domainName: %s not exist in any account", domainName)
}

// findExpireDomainsByAccount will return expire domain in specific day and account
// expireDomains map[domainName]willExpireDay
func (d *DomainClient) findExpireDomainsByAccount(expireDay int) (expireDomains expireDomainsInfo) {
	return d.getExpireDomains(expireDay)
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
	if expireDay != -1 {
		LoggerNoT.Printf("domain %s found in %s account, expire in %d days\n", domainName, accountName, expireDay)
		return
	}
	LoggerNoT.Printf("no %s found in all account", domainName)

}

func ListRegisteredDomainByAccount(i IDomainClient) {
	recordRegisterDomains := i.listRegisteredDomainByAccount()
	if len(recordRegisterDomains) > 0 {
		LoggerNoT.Printf("account %s exist registed domain\n", i.getAccountName())
		for d, _ := range recordRegisterDomains {
			LoggerNoT.Printf("domain: %s\n", d)
		}
	}
	LoggerNoT.Printf("Total count: %d\n", len(recordRegisterDomains))
}

// FindExpireDomainsByAccount will print expire domain in account
// alitool check  domain -a AccountName -e 100
func FindExpireDomainsByAccount(i IDomainClient, expireDay int) {
	expireDomains := i.findExpireDomainsByAccount(expireDay)
	if len(expireDomains) > 0 {
		for d, e := range expireDomains {
			LoggerNoT.Printf("account %s domain %s will expire in %d\n", i.getAccountName(), d, e)
		}
		return
	}
	LoggerNoT.Printf("account %s no expire domain in %d days\n", i.getAccountName(), expireDay)
}

// FindExpireDomainsInAllAccounts will print all expire domains in every account
// alitool check  domain -A -e 100
func FindExpireDomainsInAllAccounts(expireDay int) {
	initAllDomainClient()
	IDomainClients := getDomainClients()
	logrus.Infof("%#v", IDomainClients)
	for _, v := range IDomainClients {
		FindExpireDomainsByAccount(v, expireDay)
	}
}
