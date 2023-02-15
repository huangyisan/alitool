package strategy

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dcdn"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
)

type IServiceClient interface {
	NewClient(string, string, string) (interface{}, error)
}

type DomainClient struct {
}

func (*DomainClient) NewClient(regionId, accessKeyId, accessKeySecret string) (interface{}, error) {
	return domain.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
}

type DnsClient struct {
}

func (*DnsClient) NewClient(regionId, accessKeyId, accessKeySecret string) (interface{}, error) {
	return alidns.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
}

type DcdnClient struct {
}

func (*DcdnClient) NewClient(regionId, accessKeyId, accessKeySecret string) (interface{}, error) {
	return dcdn.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
}

type FinanceClient struct {
}

func (*FinanceClient) NewClient(regionId, accessKeyId, accessKeySecret string) (interface{}, error) {
	return bssopenapi.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
}

type Operator struct {
	iServiceClient IServiceClient
}

func (op *Operator) SetServiceClient(iac IServiceClient) {
	op.iServiceClient = iac
}

func (op *Operator) NewClient(regionId, accessKeyId, accessKeySecret string) (interface{}, error) {
	return op.iServiceClient.NewClient(regionId, accessKeyId, accessKeySecret)
}
