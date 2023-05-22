package finance

import (
	"alitool/internal/pkg/common"
	. "alitool/internal/pkg/mylog"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

func (f *FinanceClient) getQueryAccountBillResponse(month string) (response *bssopenapi.QueryAccountBillResponse) {
	request := bssopenapi.CreateQueryAccountBillRequest()

	request.Scheme = "https"

	request.BillingCycle = month

	res, err := f.I.QueryAccountBill(request)
	if err != nil {
		LoggerNoT.Fatalln("You are not authorized to operate this resource, or this API does not support RAM.")
	}
	return res
}

func (f *FinanceClient) getQueryAccountBalanceRequest() (response *bssopenapi.QueryAccountBalanceResponse) {
	request := bssopenapi.CreateQueryAccountBalanceRequest()
	request.Scheme = "https"
	res, err := f.I.QueryAccountBalance(request)
	if err != nil {
		LoggerNoT.Fatalln("You are not authorized to operate this resource, or this API does not support RAM.")
	}
	return res
}

func (f *FinanceClient) getLastMonthPaymentAmount() float64 {
	res := f.getQueryAccountBillResponse(common.GetLastMonth())
	if res.Data.Items.Item[0].PaymentAmount != 0 {
		return res.Data.Items.Item[0].PaymentAmount
	}
	return res.Data.Items.Item[0].PretaxAmount
}

// getBalance will return account's cash amount
func (f *FinanceClient) getAvailableAmount() string {
	res := f.getQueryAccountBalanceRequest()
	return res.Data.AvailableAmount
}
