package account

import (
	"alitool/internal/pkg/test"
	"testing"
)

func init() {
	test.GetEnv()
	InitAccount()

}
func Test_GetAccount(t *testing.T) {
	//fmt.Printf("%#v", accounts)

	GetAccount("ali_account_01")
}
