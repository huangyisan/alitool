package dns

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/strategy"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

type DnsClient struct {
	ac *alidns.Client
}

type options struct {
	regionId string
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithRegionId(regionId string) Option {
	return optionFunc(func(o *options) {
		o.regionId = regionId
	})
}

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
