package finance

import (
	"alitool/internal/pkg/common"
	. "alitool/internal/pkg/mylog"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

func (f *FinanceClient) getQueryAccountBillResponse() (response *bssopenapi.QueryAccountBillResponse) {
	request := bssopenapi.CreateQueryAccountBillRequest()

	request.Scheme = "https"

	request.BillingCycle = common.GetLastMonth()

	res, err := f.I.QueryAccountBill(request)
	if err != nil {
		LoggerNoT.Println("You are not authorized to operate this resource, or this API does not support RAM.")
	}
	return res
}
