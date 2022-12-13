package domain

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/common"
	"alitool/internal/pkg/strategy"
	"fmt"
	dm "github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
)

type DomainClient struct {
	AccountName string
	RegionId    string
	I           iDomainClient
}

type iDomainClient interface {
	QueryDomainByDomainName(request *dm.QueryDomainByDomainNameRequest) (response *dm.QueryDomainByDomainNameResponse, err error)
	QueryDomainList(request *dm.QueryDomainListRequest) (response *dm.QueryDomainListResponse, err error)
}

type IDomainClient interface {
	getAccountName() string
	listRegisteredDomainByAccount() recordRegisterDomains
	isDomainInAccount(string) bool
	getExpireDomains(int) map[string]int
	findExpireDomainRefAccount(string) (string, int)
	findExpireDomainsByAccount(IDomainClient, int) map[string]int
}

var domainClients = make([]IDomainClient, 0)

// newDomainClient will return a domain client
func newDomainClient(accountName, regionId, accessKeyId, accessKeySecret string) IDomainClient {
	op := strategy.Operator{}
	op.SetServiceClient(&strategy.DomainClient{})
	c, err := op.NewClient(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		return nil
	}
	dc, ok := c.(*dm.Client)
	if !ok {
		fmt.Println("not domain.client")
		return nil
	}
	return &DomainClient{
		accountName,
		regionId,
		dc,
	}
}

// InitDomainClient to init domain client
func InitDomainClient(accountName, regionId string) IDomainClient {
	a, ok := account.GetAccount(accountName)
	if !ok {
		return nil
	}
	dc := newDomainClient(accountName, regionId, a.GetAccessKeyId(), a.GetAccessKeySecret())
	return dc
}

// initAllDomainClient will init all DnsClient from .alitool.yaml
func initAllDomainClient() {
	accounts := account.GetAccountMap()
	for k, _ := range accounts {
		domainClients = append(domainClients, InitDomainClient(k, common.DefaultRegionId))
	}
}

func getDomainClients() []IDomainClient {
	return domainClients
}

func (d *DomainClient) getAccountName() string {
	return d.AccountName
}
