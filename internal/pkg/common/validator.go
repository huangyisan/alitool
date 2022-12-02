package common

import (
	"alitool/internal/ali/account"
	"github.com/dchest/validator"
)

func IsValidDomain(domainName string) bool {
	return validator.IsValidDomain(domainName)
}

func IsExistAccount(accountName string) bool {
	_, ok := account.GetAccount(accountName)
	if ok {
		return true
	}
	return false
}
