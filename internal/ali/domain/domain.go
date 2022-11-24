package domain

import (
	"alitool/internal/pkg/strategy"
	"fmt"
	dm "github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
	"github.com/spf13/viper"
	"strings"
)

type DomainClient struct {
	//IdomainClient
	dc *dm.Client
}

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
func initDomainClient() *DomainClient {
	regionId := viper.GetString("regionId")
	accessKeyId := viper.GetString("accessKeyId")
	accessKeySecret := viper.GetString("accessKeySecret")
	dc := NewDomainClient(regionId, accessKeyId, accessKeySecret)
	return dc
}

// domainSuffix will return domain suffix, such as www.baidu.com will return baidu.com
func domainSuffix(domainName string) string {
	dn := strings.Split(domainName, ".")
	return strings.Join(dn[len(dn)-2:], ".")
}
