package dns

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/common"
	"alitool/internal/pkg/strategy"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

type DnsClient struct {
	ac *alidns.Client
}

var dnsClients = make(map[string]*DnsClient)

// NewDnsClient return DnsClient
func NewDnsClient(regionId, accessKeyId, accessKeySecret string) *DnsClient {
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
	}
}

// initDnsClient will execute NewDnsClient to make a new DnsClient
func initDnsClient(accountName, regionId string) *DnsClient {
	a, ok := account.GetAccount(accountName)
	if !ok {
		return nil
	}
	dc := NewDnsClient(regionId, a.GetAccessKeyId(), a.GetAccessKeySecret())
	return dc
}

// InitAllDnsClient will init all DnsClient from .alitool.yaml
func InitAllDnsClient() {
	accounts := account.GetAccountMap()
	for k, _ := range accounts {
		initDnsClient(k, common.DefaultRegionId)
		dnsClients[k] = initDnsClient(k, common.DefaultRegionId)
	}
}

func GetDnsClients() map[string]*DnsClient {
	return dnsClients
}
