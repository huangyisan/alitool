package dcdn

import (
	"os"
	"testing"
)

func Test_getHttpsDomainListResponse(t *testing.T) {
	a := newDcdnClient("account01", "cn-shanghai", os.Getenv("ALITOOL_ACCESSKEYID"), os.Getenv("ALITOOL_ACCESSKEYSECRET"))
	res, keys := a.listAllHttpsDomains()
	for _, v := range keys {
		t.Logf("%s\n\t%+v\n", v, res[v])
	}
}
