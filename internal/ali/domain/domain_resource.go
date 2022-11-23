package domain

import "fmt"

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

func DoGetDomainResource() {
	dc := initDomainClient()
	dc.getDomainResource()
}
