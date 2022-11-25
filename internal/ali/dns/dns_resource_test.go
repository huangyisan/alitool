package dns

import (
	"testing"
)

//func init() {
//	test.GetEnv()
//}

//func Test_getDNSRecords(t *testing.T) {
//	c := initDnsClient()
//	c.getDNSRecords("xiaozhumao.com")
//	//c.getDNSRecord("dl.zhibosha.com")
//}

func Test_DoIsDNSExist(t *testing.T) {
	DoIsDNSExist("ali_account_01", "ccc", "cc")
}
