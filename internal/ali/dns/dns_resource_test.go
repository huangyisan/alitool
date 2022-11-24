package dns

import (
	"alitool/internal/pkg/test"
	"testing"
)

func init() {
	test.GetEnv()
}

func Test_getDNSRecords(t *testing.T) {
	c := initDnsClient()
	c.getDNSRecords("xiaozhumao.com")
	//c.getDNSRecord("dl.zhibosha.com")
}

func Test_DoIsDNSExist(t *testing.T) {
	DoIsDNSExist("zhibosha.com")
}
