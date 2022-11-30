package dns

import "testing"

func Test_ListDnsByAccount(t *testing.T) {
	setup()
	ListDnsByAccount("ali_account_01")
}
