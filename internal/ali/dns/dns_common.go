package dns

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

func (d *DnsClient) DescribeDomainRecordsViaA(domainName string, subDomain chan string) error {
	request := makeRequest("A", domainName)
	response, err := d.ac.DescribeDomainRecords(request)
	if err != nil {
		return err
	}
	record := response.DomainRecords.Record
	for _, v := range record {
		if v.RR != "@" {
			subDomain <- fmt.Sprintf("%s.%s", v.RR, v.DomainName)
		}
	}
	return nil
}

func (d *DnsClient) DescribeDomainRecordsViaCNAME(domainName string, subDomain chan string) error {
	request := makeRequest("CNAME", domainName)
	response, err := d.ac.DescribeDomainRecords(request)
	if err != nil {
		return err
	}
	record := response.DomainRecords.Record
	for _, v := range record {
		if v.RR != "@" {
			subDomain <- fmt.Sprintf("%s.%s", v.RR, v.DomainName)
		}
	}
	return nil
}

func (d *DnsClient) DescribeDomains() (hasRecordDomains []string) {
	request := alidns.CreateDescribeDomainsRequest()
	request.Scheme = "https"
	request.PageSize = "100"
	response, err := d.ac.DescribeDomains(request)
	if err != nil {
		fmt.Print(err.Error())
	}

	// 遍历结果
	for _, v := range response.Domains.Domain {
		if v.RecordCount != 0 {
			hasRecordDomains = append(hasRecordDomains, v.DomainName)
		}
	}
	return
}

func makeRequest(dnsType, domainName string) (request *alidns.DescribeDomainRecordsRequest) {
	request = alidns.CreateDescribeDomainRecordsRequest()
	request.Scheme = "https"
	request.Status = "Enable"
	request.DomainName = domainName
	request.Type = dnsType
	return
}

//func (ad accountDomain) GetSubDomainsByAccount(ac Account, subDomain chan string) error {
//	var err error
//	for _, v := range ad.domains {
//		err = ac.DescribeDomainRecordsViaA(v, subDomain)
//		err = ac.DescribeDomainRecordsViaCNAME(v, subDomain)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
