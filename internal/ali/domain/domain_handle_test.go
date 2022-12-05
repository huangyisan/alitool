package domain

import (
	"alitool/internal/pkg/common"
	"fmt"
	"testing"
)

func TestDomainClient_checkDomainStatus(t *testing.T) {
	//dc := NewDomainClient("cn-shanghai", "", "")
	//dc.getExpireDomains(300)
}

func Test_spilit(t *testing.T) {
	word := "baidu.com"
	fmt.Println(common.DomainSuffix(word))
}

func Test_findExpireDomainsInAllAccounts(t *testing.T) {
	setup()
	a := findExpireDomainsInAllAccounts(300)
	for account, v := range a {
		fmt.Println(account)
		for domain, expireday := range v {
			fmt.Println(domain, expireday)
		}

	}
}

func Test_FindExpireDomainRefAccount(t *testing.T) {
	setup()
	FindExpireDomainRefAccount("xiaozhumao.com")
}
