package dns

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

// _common.go file will return origin resource information

// DescribeDomainRecordsViaA print domain all A records
func (d *DnsClient) DescribeDomainRecordsViaA(domainName string) []alidns.Record {
	request := makeRequest("A", domainName)
	response, err := d.ac.DescribeDomainRecords(request)
	if err != nil {
		return nil
	}
	return response.DomainRecords.Record

}

// DescribeDomainRecordsViaCNAME print domain all CNAME records
func (d *DnsClient) DescribeDomainRecordsViaCNAME(domainName string) []alidns.Record {
	request := makeRequest("CNAME", domainName)
	response, err := d.ac.DescribeDomainRecords(request)
	if err != nil {
		return nil
	}
	return response.DomainRecords.Record
}

// ListDomains return all domains in ali account
func (d *DnsClient) ListDomains() (hasRecordDomains map[string]struct{}) {
	hasRecordDomains = make(map[string]struct{})
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
			hasRecordDomains[v.DomainName] = struct{}{}
		}
	}
	return hasRecordDomains
}

// makeRequest encapsulate request
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
