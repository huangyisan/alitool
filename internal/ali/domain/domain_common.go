package domain

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	dm "github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
)

// _common.go file will return origin resource information

// queryDomainByDomainNameInfo will return domain all information
func (dc *DomainClient) queryDomainByDomainNameInfo(domainName string) *dm.QueryDomainByDomainNameResponse {
	request := dm.CreateQueryDomainByDomainNameRequest()
	request.Scheme = "https"
	request.DomainName = domainName
	response, err := dc.dc.QueryDomainByDomainName(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	return response
}

// getAllDomains will return all domains in ali account
func (dc *DomainClient) getAllDomains() (responses []*dm.QueryDomainListResponse) {
	var pageStartNumber = 1
	nextFlag := true
	request := dm.CreateQueryDomainListRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(20)

	for nextFlag {
		request.PageNum = requests.NewInteger(pageStartNumber)
		response, err := dc.dc.QueryDomainList(request)
		if err != nil {
			fmt.Println(err.Error())
		}
		responses = append(responses, response)
		if response.NextPage == false {
			nextFlag = false
		}
		pageStartNumber += 1
	}
	return responses
}
