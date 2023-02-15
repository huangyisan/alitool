package finance

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/common"
	"alitool/internal/pkg/strategy"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

type FinanceClient struct {
	AccountName string
	RegionId    string
	I           iFinanceClient
}

type iFinanceClient interface {
	QueryAccountBill(request *bssopenapi.QueryAccountBillRequest) (response *bssopenapi.QueryAccountBillResponse, err error)
	//QueryDomainByDomainName(request *dm.QueryDomainByDomainNameRequest) (response *dm.QueryDomainByDomainNameResponse, err error)
	//QueryDomainList(request *dm.QueryDomainListRequest) (response *dm.QueryDomainListResponse, err error)
}

type IFinanceClient interface {
	getLastMonthPaymentAmount() float64
	getAccountName() string
}

// newFinanceClient will return a finance client
func newFinanceClient(accountName, regionId, accessKeyId, accessKeySecret string) IFinanceClient {
	op := strategy.Operator{}
	op.SetServiceClient(&strategy.FinanceClient{})
	c, err := op.NewClient(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		return nil
	}
	bss, ok := c.(*bssopenapi.Client)
	if !ok {
		fmt.Println("not finance.client")
		return nil
	}
	return &FinanceClient{
		accountName,
		regionId,
		bss,
	}
}

// InitFinanceClient to init finance client
func InitFinanceClient(accountName, regionId string) IFinanceClient {
	a, ok := account.GetAccount(accountName)
	if !ok {
		return nil
	}
	dc := newFinanceClient(accountName, regionId, a.GetAccessKeyId(), a.GetAccessKeySecret())
	return dc
}

var financeClients = make([]IFinanceClient, 0)

// initAllDomainClient will init all DnsClient from .alitool.yaml
func initAllDomainClient() {
	accounts := account.GetAccountMap()
	for k, _ := range accounts {
		financeClients = append(financeClients, InitFinanceClient(k, common.DefaultRegionId))
	}
}

func getFinanceClients() []IFinanceClient {
	return financeClients
}

// getAccountName will return DnsClient's AccountName
func (f *FinanceClient) getAccountName() string {
	return f.AccountName
}
