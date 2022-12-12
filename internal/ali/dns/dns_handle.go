package dns

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

type recordDnsDomains map[string]struct{}

// getAllDnsDomains will return all dns domains in ali account
func (d *DnsClient) getAllDnsDomains() (hasRecordDomains recordDnsDomains) {
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
		response, err := d.I.DescribeDomains(request)
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
