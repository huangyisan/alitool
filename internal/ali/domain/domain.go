package domain

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/common"
	"alitool/internal/pkg/strategy"
	"fmt"
	dm "github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
)

type DomainClient struct {
	dc *dm.Client
}

var domainClients = make(map[string]*DomainClient)

// NewDomainClient will return a domain client
func NewDomainClient(regionId, accessKeyId, accessKeySecret string) *DomainClient {
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
		dc,
	}
}

// initDomainClient to init domain client
func initDomainClient(accountName, regionId string) *DomainClient {
	a, ok := account.GetAccount(accountName)
	if !ok {
		return nil
	}
	dc := NewDomainClient(regionId, a.GetAccessKeyId(), a.GetAccessKeySecret())
	return dc
}

// InitAllDomainClient will init all DnsClient from .alitool.yaml
func InitAllDomainClient() {
	accounts := account.GetAccountMap()
	for k, _ := range accounts {
		initDomainClient(k, common.DefaultRegionId)
		domainClients[k] = initDomainClient(k, common.DefaultRegionId)
	}
}

func GetDomainClients() map[string]*DomainClient {
	return domainClients
}
