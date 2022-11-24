package domain

import "fmt"

// getDomainResource print all the domains register in ali account
func (dc *DomainClient) getDomainResource() {

	domainSlice := make([]string, 0)

	for _, dms := range dc.getAllDomains() {
		for _, dm := range dms.Data.Domain {
			domainSlice = append(domainSlice, dm.DomainName)
		}
	}

	ldc := len(domainSlice)
	if ldc > 0 {
		fmt.Printf("exist domain count %d:\n", ldc)
		for _, dms := range dc.getAllDomains() {
			for _, dm := range dms.Data.Domain {
				fmt.Printf("%s\n", dm.DomainName)
			}
		}
		return
	}
	fmt.Printf("Domains count %d:\n", ldc)
}

// DoGetDomainResource execute domainResource function
func DoGetDomainResource() {
	dc := initDomainClient()
	dc.getDomainResource()
}
