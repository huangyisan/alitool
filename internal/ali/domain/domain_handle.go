package domain

import (
	"alitool/internal/pkg/common"
	. "alitool/internal/pkg/mylog"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	dm "github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
)

type recordRegisterDomains map[string]struct{}

// getRegisteredDomainResponse will return domain response
func (d *DomainClient) getRegisteredDomainResponse(domainName string) *dm.QueryDomainByDomainNameResponse {
	request := dm.CreateQueryDomainByDomainNameRequest()
	request.Scheme = "https"
	request.DomainName = domainName
	response, err := d.I.QueryDomainByDomainName(request)
	if err != nil {
		LoggerNoT.Println(err.Error())
	}
	return response
}

// getAllRegisteredDomainsResponse will return response
func (d *DomainClient) getAllRegisteredDomainsResponse() (response []*dm.QueryDomainListResponse) {
	var pageStartNumber = 1
	var nextFlag = true
	var pageSize = 20
	request := dm.CreateQueryDomainListRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(pageSize)

	for nextFlag {
		request.PageNum = requests.NewInteger(pageStartNumber)
		res, err := d.I.QueryDomainList(request)
		if err != nil {
			LoggerNoT.Println(err.Error())
			return nil
		}
		// literal results
		response = append(response, res)

		if res.NextPage == false {
			nextFlag = false
		}
		pageStartNumber += 1
	}
	return response
}

// getAllRegisteredDomains will return all domains in ali account
func (d *DomainClient) getAllRegisteredDomains() (hasRecordDomains recordRegisterDomains) {
	hasRecordDomains = make(map[string]struct{})
	allRegisteredDomainsResponse := d.getAllRegisteredDomainsResponse()
	for _, res := range allRegisteredDomainsResponse {
		for _, v := range res.Data.Domain {
			hasRecordDomains[v.DomainName] = struct{}{}
		}
	}
	return hasRecordDomains
}

// getExpireDomains will print all the domain expire day in ali account
func (d *DomainClient) getExpireDomains(expireDay int) (expireDomains expireDomainsInfo) {
	expireDomains = make(map[string]int)
	for _, dms := range d.getAllRegisteredDomainsResponse() {
		for _, _dm := range dms.Data.Domain {
			if _dm.ExpirationCurrDateDiff < expireDay {
				expireDomains[_dm.DomainName] = _dm.ExpirationCurrDateDiff
			}
		}
	}
	if len(expireDomains) > 0 {
		return expireDomains
	}
	return expireDomains
}

// getDomainExpireCurrDiff will print the specific domain will remain someday to expire
func (d *DomainClient) getDomainExpireCurrDiff(domainName string) (expireDay int) {
	_domainName := common.DomainSuffix(domainName)
	expireDay = d.getRegisteredDomainResponse(_domainName).ExpirationCurrDateDiff
	return
}
