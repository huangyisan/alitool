package ssl

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/common"
	"alitool/internal/pkg/strategy"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cas"
)

type SSLClient struct {
	AccountName string
	RegionId    string
	I           iSSLClient
}

type iSSLClient interface {
	ListUserCertificateOrder(request *cas.ListUserCertificateOrderRequest) (response *cas.ListUserCertificateOrderResponse, err error)
}

type ISSLClient interface {
	getAccountName() string
	getExpireCertByAccount(expireDay int) (certs expireCertsInfo)
	getExpireUploadCertByAccount(expireDay int) (certs expireCertsInfo)
}

// newSSLClient will return a ssl client
func newSSLClient(accountName, regionId, accessKeyId, accessKeySecret string) ISSLClient {
	op := strategy.Operator{}
	op.SetServiceClient(&strategy.SSLClient{})
	c, err := op.NewClient(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		return nil
	}
	cc, ok := c.(*cas.Client)
	if !ok {
		fmt.Println("not domain.client")
		return nil
	}
	return &SSLClient{
		accountName,
		regionId,
		cc,
	}
}

// InitSSLClient to init ssl client
func InitSSLClient(accountName, regionId string) ISSLClient {
	a, ok := account.GetAccount(accountName)
	if !ok {
		return nil
	}
	dc := newSSLClient(accountName, regionId, a.GetAccessKeyId(), a.GetAccessKeySecret())
	return dc
}

var sslClients = make([]ISSLClient, 0)

// initAllSSLClient will init all SSLClient from .alitool.yaml
func initAllSSLClient() {
	accounts := account.GetAccountMap()
	if len(sslClients) == 0 {
		for k, _ := range accounts {
			sslClients = append(sslClients, InitSSLClient(k, common.DefaultRegionId))
		}
	}
}

func getSSLClients() []ISSLClient {
	return sslClients
}

// getAccountName will return Finance's AccountName
func (s *SSLClient) getAccountName() string {
	return s.AccountName
}
