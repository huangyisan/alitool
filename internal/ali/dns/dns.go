package dns

import (
	"alitool/internal/pkg/strategy"
	"alitool/internal/pkg/test"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/spf13/viper"
)

type DnsClient struct {
	ac *alidns.Client
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
func initDnsClient() *DnsClient {
	test.GetEnv()
	regionId := viper.GetString("regionId")
	accessKeyId := viper.GetString("accessKeyId")
	accessKeySecret := viper.GetString("accessKeySecret")
	dc := NewDnsClient(regionId, accessKeyId, accessKeySecret)
	return dc
}
