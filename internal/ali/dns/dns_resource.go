package dns

import "fmt"

func (d *DnsClient) getDNSRecords(domainName string) {
	cnameRecords := d.DescribeDomainRecordsViaCNAME(domainName)
	aRecords := d.DescribeDomainRecordsViaA(domainName)
	if cnameRecords != nil || aRecords != nil {
		for _, v := range cnameRecords {
			fmt.Printf("%s CNAME %s\n", v.RR, v.Value)
		}
		for _, v := range aRecords {
			fmt.Printf("%s A %s\n", v.RR, v.Value)
		}
		return
	}

	fmt.Printf("%s not in this account or other type", domainName)
}

// IsDomainDNSExist judge domain dns record in ali account
func (d *DnsClient) isDNSExist(domainName string) {
	_, ok := d.ListDomains()[domainName]
	if ok {
		fmt.Printf("exist domain: %s", domainName)
		return
	}
	fmt.Printf("not exist domain: %s", domainName)
}

func DoIsDNSExist(accountName, domainName, regionId string) {
	dc := initDnsClient(accountName, regionId)
	dc.isDNSExist(domainName)
}
