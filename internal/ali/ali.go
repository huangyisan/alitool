package ali

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)

type AliClient struct {
	sdk.Client
}

func NewAliyunClient(regionID, accessKeyID, accessKeySecret string) *AliClient {
	ac, err := sdk.NewClientWithAccessKey(regionID, accessKeyID, accessKeySecret)
	if err != nil {
		panic(err)
	}
	return &AliClient{
		*ac,
	}

}
