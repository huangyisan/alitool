package account

import (
	. "alitool/internal/pkg/mylog"
)

// ListAccount will print all AliAccount
func ListAccount() {
	LoggerNoT.Infof("Account List:\n\n")
	for k, _ := range getAccountMap() {
		LoggerNoT.Infof("%s\n", k)
	}
}

// ListSubAccount will print all AliSubAccount
func ListSubAccount() {
	LoggerNoT.Infof("SubAccount List:\n\n")
	for k, v := range getAccountMap() {
		LoggerNoT.Infof("%s subAccount: %s\n", k, v["subAccount"])
	}
}
