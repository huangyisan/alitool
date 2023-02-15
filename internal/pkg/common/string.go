package common

import (
	"strings"
	"time"
)

// DomainSuffix will return domain suffix, such as www.baidu.com will return baidu.com
func DomainSuffix(domainName string) string {
	dn := strings.Split(domainName, ".")
	if len(dn) == 1 {
		return strings.Join(dn, "")
	}
	return strings.Join(dn[len(dn)-2:], ".")
}

// GetLastMonth will return last month, such as "2022-12"
func GetLastMonth() string {
	lastMonth := time.Now().AddDate(0, -1, 0)
	return lastMonth.Format("2006-01")
}
