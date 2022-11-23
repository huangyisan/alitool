package dns

import (
	"alitool/internal/pkg/strategy"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/spf13/viper"
)

type DnsClient struct {
	//IdomainClient
	ac *alidns.Client
}

func NewDnsClient(regionId, accessKeyId, accessKeySecret string) *DnsClient {
	op := strategy.Operator{}
	op.SetServiceClient(&strategy.DomainClient{})
	c, err := op.NewClient(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		return nil
	}
	dc, ok := c.(*alidns.Client)
	if !ok {
		fmt.Println("not domain.client")
		return nil
	}
	return &DnsClient{
		dc,
	}
}

func initDnsClient() *DnsClient {
	regionId := viper.GetString("regionId")
	accessKeyId := viper.GetString("accessKeyId")
	accessKeySecret := viper.GetString("accessKeySecret")
	dc := NewDnsClient(regionId, accessKeyId, accessKeySecret)
	return dc
}
