package dns

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/common"
	"alitool/internal/pkg/strategy"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

var (
	_ IDNSClient = (*DnsClient)(nil)
)

type DnsClient struct {
	AccountName string
	RegionId    string
	I           iDnsClient
}

// var dnsClients = make(map[string]*DnsClient)
var dnsClients = make([]IDNSClient, 0)

type iDnsClient interface {
	DescribeDomains(request *alidns.DescribeDomainsRequest) (response *alidns.DescribeDomainsResponse, err error)
}

type IDNSClient interface {
	getAccountName() string
	listDnsByAccount() recordDnsDomains
	isDnsInAccount(string) bool
}

// newDnsClient return DnsClient
func newDnsClient(accountName, regionId, accessKeyId, accessKeySecret string) IDNSClient {
	op := strategy.Operator{}
	op.SetServiceClient(&strategy.DnsClient{})
	c, err := op.NewClient(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		return nil
	}
	dc, ok := c.(*alidns.Client)
	if !ok {
		fmt.Println("not alidns.client")
		return nil
	}
	return &DnsClient{
		AccountName: accountName,
		RegionId:    regionId,
		I:           dc,
	}
}

// InitDnsClient will execute newDnsClient to make a new DnsClient
func InitDnsClient(accountName, regionId string) IDNSClient {
	a, ok := account.GetAccount(accountName)
	if !ok {
		return nil
	}
	dc := newDnsClient(accountName, regionId, a.GetAccessKeyId(), a.GetAccessKeySecret())
	return dc
}

// initAllDnsClients will init all DnsClient from .alitool.yaml
func initAllDnsClients() {
	accounts := account.GetAccountMap()
	for k, _ := range accounts {
		dnsClients = append(dnsClients, InitDnsClient(k, common.DefaultRegionId))
	}
}

// getDnsClients will return dnsClients
func getDnsClients() []IDNSClient {
	return dnsClients
}

// getAccountName will return DnsClient's AccountName
func (d *DnsClient) getAccountName() string {
	return d.AccountName
}
