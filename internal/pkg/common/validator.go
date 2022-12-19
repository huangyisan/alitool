package common

import (
	"github.com/dchest/validator"
)

func IsValidDomain(domainName string) bool {
	return validator.IsValidDomain(domainName)
}
