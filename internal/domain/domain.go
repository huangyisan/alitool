package domain

import (
	"alitool/internal/pkg/strategy"
	"fmt"
	dm "github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
)

type DomainClient struct {
	//IdomainClient
	dc *dm.Client
}

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
