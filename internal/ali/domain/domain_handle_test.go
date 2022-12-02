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
