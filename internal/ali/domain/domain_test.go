package domain

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/test"
)

//func TestDomainClie(t *testing.T) {
//	newDomainClient()
//}

func setup() {
	test.GetEnv()
	account.InitAccount()
	InitAllDomainClient()
}
