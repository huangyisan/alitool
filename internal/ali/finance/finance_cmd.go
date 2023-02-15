package finance

import (
	"alitool/internal/pkg/common"
	. "alitool/internal/pkg/mylog"
)

func ListLastMonthPaymentAmountByAccount(i IFinanceClient) {
	paymentAmount := i.getLastMonthPaymentAmount()
	LoggerNoT.Printf("%s %s payment amount: %f\n", common.GetLastMonth(), i.getAccountName(), paymentAmount)
}

func ListLastMonthPaymentAmountInAllAccounts() {
	initAllDomainClient()
	IFinanceClients := getFinanceClients()
	for _, v := range IFinanceClients {
		ListLastMonthPaymentAmountByAccount(v)
	}
}
