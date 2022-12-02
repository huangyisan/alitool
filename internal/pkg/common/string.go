package common

import (
	"strings"
)

// DomainSuffix will return domain suffix, such as www.baidu.com will return baidu.com
func DomainSuffix(domainName string) string {
	dn := strings.Split(domainName, ".")
	if len(dn) == 1 {
		return strings.Join(dn, "")
	}
	return strings.Join(dn[len(dn)-2:], ".")
}
