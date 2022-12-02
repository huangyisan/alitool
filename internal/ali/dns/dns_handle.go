package dns

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

type RecordDomains map[string]struct{}

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

// GetAllDomains return all dns domains in ali account
func (d *DnsClient) GetAllDomains() (hasRecordDomains RecordDomains) {
	hasRecordDomains = make(map[string]struct{})

	var pageStartNumber = 1
	var totalCount int64
	var pageSize = 20
	nextFlag := true
	request := alidns.CreateDescribeDomainsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(pageSize)

	for nextFlag {
		request.PageNumber = requests.NewInteger(pageStartNumber)
		response, err := d.ac.DescribeDomains(request)
		// get total count
		totalCount = response.TotalCount
		if err != nil {
			fmt.Print(err.Error())
		}
		// literal results
		for _, v := range response.Domains.Domain {
			hasRecordDomains[v.DomainName] = struct{}{}
		}
		if pageStartNumber*pageSize >= int(totalCount) {
			nextFlag = false
		}
		pageStartNumber += 1
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
