package domain

import "fmt"

type day int

func (dc *DomainClient) GetDomainExpireCurrDiff(domainName string) {
	days := dc.queryDomainByDomainNameInfo(domainName).ExpirationCurrDateDiff
	fmt.Printf("Domain: %s will expire in %d days\n", domainName, days)
}

func (dc *DomainClient) GetExpireDomains(remainDays day) {
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
