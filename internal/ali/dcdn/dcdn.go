package dcdn

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/common"
	"alitool/internal/pkg/strategy"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dcdn"
)

type DcdnClient struct {
	AccountName string
	RegionId    string
	I           iDcdnClient
}

type iDcdnClient interface {
	DescribeDcdnHttpsDomainList(request *dcdn.DescribeDcdnHttpsDomainListRequest) (response *dcdn.DescribeDcdnHttpsDomainListResponse, err error)
	DescribeDcdnUserDomains(request *dcdn.DescribeDcdnUserDomainsRequest) (response *dcdn.DescribeDcdnUserDomainsResponse, err error)
}

type IDcdnClient interface {
	getHttpsDomainListResponse() (response []*dcdn.DescribeDcdnUserDomainsResponse)
	listAllHttpsDomains() (hasRecordHttpsDomains recordHttpsDomains, hasRecordHttpsDomainsKeys []string)
	getDcdnClientName() string
}

var dcdnClients = make([]IDcdnClient, 0)

func newDcdnClient(accountName, regionId, accessKeyId, accessKeySecret string) IDcdnClient {
	op := strategy.Operator{}
	op.SetServiceClient(&strategy.DcdnClient{})
	c, err := op.NewClient(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		return nil
	}
	dc, ok := c.(*dcdn.Client)
	if !ok {
		fmt.Println("not dcdn.client")
		return nil
	}
	return &DcdnClient{
		AccountName: accountName,
		RegionId:    regionId,
		I:           dc,
	}
}

// InitDcdnClient to init domain client
func InitDcdnClient(accountName, regionId string) IDcdnClient {
	a, ok := account.GetAccount(accountName)
	if !ok {
		return nil
	}
	dc := newDcdnClient(accountName, regionId, a.GetAccessKeyId(), a.GetAccessKeySecret())
	return dc
}

// initAllDcdnClient will init all DnsClient from .alitool.yaml
func initAllDcdnClient() {
	accounts := account.GetAccountMap()
	for k, _ := range accounts {
		dcdnClients = append(dcdnClients, InitDcdnClient(k, common.DefaultRegionId))
	}
}

func getDcdnClients() []IDcdnClient {
	return dcdnClients
}
