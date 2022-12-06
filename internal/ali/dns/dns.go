package dns

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/common"
	"alitool/internal/pkg/strategy"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

type DnsClient struct {
	ac          *alidns.Client
	accountName string
	regionId    string
}

type IDNSClient interface {
	getAccountName() string
	listDnsByAccount() recordDnsDomains
	isDnsInAccount(string) bool
}

// var dnsClients = make(map[string]*DnsClient)
var dnsClients = make([]IDNSClient, 0)

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
		dc,
		accountName,
		regionId,
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

func getDnsClients() []IDNSClient {
	return dnsClients
}

func (d *DnsClient) getAccountName() string {
	return d.accountName
}
