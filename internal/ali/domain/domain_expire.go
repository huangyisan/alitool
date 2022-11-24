package domain

import (
	"fmt"
)

// getDomainExpireCurrDiff will print the specific domain will remain someday to expire
func (dc *DomainClient) getDomainExpireCurrDiff(domainName string) {
	dm := domainSuffix(domainName)
	days := dc.queryDomainByDomainNameInfo(dm).ExpirationCurrDateDiff
	fmt.Printf("Domain: %s will expire in %d days\n", dm, days)
}

// getExpireDomains will print all the domain expire day in ali account
func (dc *DomainClient) getExpireDomains(remainDays int) {
	domainMap := make(map[string]int)
	for _, dms := range dc.getAllDomains() {
		for _, dm := range dms.Data.Domain {
			if dm.ExpirationCurrDateDiff < int(remainDays) {
				domainMap[dm.DomainName] = dm.ExpirationCurrDateDiff
			}
		}
	}
	if len(domainMap) > 1 {
		for k, v := range domainMap {
			fmt.Printf("Domain: %s will expire in %d days\n", k, v)
		}
	}
}

// DoGetExpireDomains execute expireDomains function
func DoGetExpireDomains(day int) {
	dc := initDomainClient()
	dc.getExpireDomains(day)
}

// DoGetDomainExpireCurrDiff execute domainExpireCurrDiff function
func DoGetDomainExpireCurrDiff(domainName string) {
	dc := initDomainClient()
	dc.getDomainExpireCurrDiff(domainName)
}
