package finance

import (
	"alitool/internal/pkg/common"
	. "alitool/internal/pkg/mylog"
)

func ListLastMonthPaymentAmount(i IFinanceClient) {
	paymentAmount := i.getLastMonthPaymentAmount()
	LoggerNoT.Printf("%s: payment amount: %f\n", common.GetLastMonth(), paymentAmount)
}
