package dcdn

import (
	. "alitool/internal/pkg/mylog"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dcdn"
	"sort"
)

type httpsDomainsDetail struct {
	Source []string
}

type recordHttpsDomains map[string]httpsDomainsDetail

func (d *DcdnClient) getHttpsDomainListResponse() (response []*dcdn.DescribeDcdnUserDomainsResponse) {
	var pageStartNumber = 1
	var totalCount int64
	var pageSize = 20
	nextFlag := true

	request := dcdn.CreateDescribeDcdnUserDomainsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(pageSize)
	request.DomainStatus = "online"
	for nextFlag {
		request.PageNumber = requests.NewInteger(pageStartNumber)
		res, err := d.I.DescribeDcdnUserDomains(request)
		if err != nil {
			LoggerNoT.Println("You are not authorized to operate this resource, or this API does not support RAM.")
			//LoggerNoT.Println(err.Error())
		}
		totalCount = res.TotalCount
		response = append(response, res)
		if pageStartNumber*pageSize >= int(totalCount) {
			nextFlag = false
		}
		pageStartNumber += 1
	}
	return response
}

// getAllHttpsDomains return https domains, http domain will ignore
func (d *DcdnClient) listAllHttpsDomains() (hasRecordHttpsDomains recordHttpsDomains, hasRecordHttpsDomainsKeys []string) {
	hasRecordHttpsDomains = make(recordHttpsDomains)
	hasRecordHttpsDomainsKeys = make([]string, 0)
	res := d.getHttpsDomainListResponse()
	for _, v := range res {
		for _, y := range v.Domains.PageData {
			if y.SSLProtocol == "on" {
				tmpSource := make([]string, 0)
				for _, v := range y.Sources.Source {
					tmpSource = append(tmpSource, fmt.Sprintf("%s:%d", v.Content, v.Port))
				}
				hasRecordHttpsDomains[y.DomainName] = httpsDomainsDetail{
					Source: tmpSource,
				}
			}
		}
	}
	for k, _ := range hasRecordHttpsDomains {
		hasRecordHttpsDomainsKeys = append(hasRecordHttpsDomainsKeys, k)
	}
	sort.Strings(hasRecordHttpsDomainsKeys)

	return hasRecordHttpsDomains, hasRecordHttpsDomainsKeys
}

func (d *DcdnClient) getDcdnClientName() string {
	return d.AccountName
}
