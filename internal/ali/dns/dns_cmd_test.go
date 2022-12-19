package dns

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/mylog"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func Test_listDnsByAccount(t *testing.T) {
	var d *DnsClient
	convey.Convey("Mock getAllDnsDomains", t, func() {
		patches := gomonkey.ApplyPrivateMethod(reflect.TypeOf(d), "getAllDnsDomains", func(*DnsClient) recordDnsDomains {
			return recordDnsDomains{"test.com": struct{}{}}
		})
		defer patches.Reset()
		convey.Convey("listDnsByAccount", func() {
			hasRecordDomains := d.listDnsByAccount()
			convey.So(hasRecordDomains, convey.ShouldResemble, recordDnsDomains{"test.com": struct{}{}})
		})
	})
}

func Test_isDnsInAccount(t *testing.T) {
	var d *DnsClient
	trueCaseDomain := "test.com"
	falseCaseDomain := "false.com"
	convey.Convey("Mock listDnsByAccount", t, func() {
		patches := gomonkey.ApplyPrivateMethod(reflect.TypeOf(d), "listDnsByAccount", func(*DnsClient) recordDnsDomains {
			return recordDnsDomains{"test.com": struct{}{}}
		})
		defer patches.Reset()

		convey.Convey("Give trueCaseDomain", func() {
			res := d.isDnsInAccount(trueCaseDomain)
			convey.So(res, convey.ShouldEqual, true)
		})

		convey.Convey("Give falseCaseDomain", func() {
			res := d.isDnsInAccount(falseCaseDomain)
			convey.So(res, convey.ShouldEqual, false)
		})
	})
}

var mockAccountName = "mockAccountName"
var mockDnsRecord = "test.com"

type mockIDNSClient struct {
	dnsInAccountFlag bool
}

func (m *mockIDNSClient) getAccountName() string {
	return mockAccountName
}

func (m *mockIDNSClient) listDnsByAccount() recordDnsDomains {
	return map[string]struct{}{mockDnsRecord: {}}
}

func (m *mockIDNSClient) isDnsInAccount(string) bool {
	return m.dnsInAccountFlag
}

func (m *mockIDNSClient) setDnsInAccountFlag(flag bool) {
	m.dnsInAccountFlag = flag
}

func TestListDnsByAccount(t *testing.T) {
	convey.Convey("Mock account.IsExistAccount false", t, func() {
		patches := gomonkey.ApplyFunc(account.IsExistAccount, func(accountName string) bool {
			return false
		})
		defer patches.Reset()
		var mockIDNSClient mockIDNSClient
		var a = func() {
			func(i IDNSClient) {
				ListDnsByAccount(i)
			}(&mockIDNSClient)

		}
		res := mylog.CaptureLogInfo(a)
		convey.So(res, convey.ShouldContainSubstring, "is right?")
	})

	convey.Convey("Mock account.IsExistAccount true", t, func() {
		patches := gomonkey.ApplyFunc(account.IsExistAccount, func(accountName string) bool {
			return true
		})
		defer patches.Reset()
		var mockIDNSClient mockIDNSClient
		var a = func() {
			func(i IDNSClient) {
				ListDnsByAccount(i)
			}(&mockIDNSClient)

		}
		res := mylog.CaptureLogInfo(a)
		convey.So(res, convey.ShouldContainSubstring, mockAccountName)
		convey.So(res, convey.ShouldContainSubstring, mockDnsRecord)
	})
}

func TestIsDnsInAccount(t *testing.T) {
	convey.Convey("Mock account.IsExistAccount return false", t, func() {
		patches := gomonkey.ApplyFunc(account.IsExistAccount, func(accountName string) bool {
			return false
		})
		defer patches.Reset()

		var i mockIDNSClient
		domainName := "test.com"
		var a = func() {
			func(i IDNSClient, domainName string) {
				IsDnsInAccount(i, domainName)
			}(&i, domainName)
		}
		res := mylog.CaptureLogInfo(a)
		convey.So(res, convey.ShouldContainSubstring, "invalid")

	})
	convey.Convey("Mock account.IsExistAccount return true", t, func() {
		patches := gomonkey.ApplyFunc(account.IsExistAccount, func(accountName string) bool {
			return true
		})
		defer patches.Reset()
		convey.Convey("Give invalid domainName", func() {
			var i mockIDNSClient
			domainName := "com"
			var a = func() {
				func(i IDNSClient, domainName string) {
					IsDnsInAccount(i, domainName)
				}(&i, domainName)
			}
			res := mylog.CaptureLogInfo(a)
			convey.So(res, convey.ShouldContainSubstring, "invalid")

		})

		convey.Convey("Give valid domainName", func() {
			convey.Convey("dns in account", func() {
				//var i mockIDNSClient
				domainName := "test.com"

				var a = func() {
					i := mockIDNSClient{}
					i.setDnsInAccountFlag(true)
					func(i IDNSClient, domainName string) {
						IsDnsInAccount(i, domainName)
					}(&i, domainName)
				}
				res := mylog.CaptureLogInfo(a)
				convey.So(res, convey.ShouldNotContainSubstring, "not")

			})

			convey.Convey("dns not in account", func() {
				var i mockIDNSClient
				domainName := "test.com"
				i.setDnsInAccountFlag(false)
				var a = func() {
					func(i IDNSClient, domainName string) {
						IsDnsInAccount(i, domainName)
					}(&i, domainName)
				}
				res := mylog.CaptureLogInfo(a)
				convey.So(res, convey.ShouldContainSubstring, "not")
			})
		})
	})
}

func TestFindDnsInAccount(t *testing.T) {
	convey.Convey("Mock getDnsClients []IDNSClient{} is empty", t, func() {
		patches := gomonkey.ApplyFunc(getDnsClients, func() []IDNSClient {
			return []IDNSClient{}
		})
		defer patches.Reset()
		var a = func() {
			FindDnsInAccount("test.com")
		}
		res := mylog.CaptureLogInfo(a)
		convey.So(res, convey.ShouldContainSubstring, "not")
	})

	convey.Convey("Mock getDnsClients []IDNSClient{} is not empty", t, func() {
		var i mockIDNSClient
		var c mockIDNSClient
		patches := gomonkey.ApplyFunc(getDnsClients, func() []IDNSClient {
			return []IDNSClient{&i, &c}
		})
		defer patches.Reset()
		convey.Convey("isDnsInAccount return true", func() {
			i.setDnsInAccountFlag(true)
			var a = func() {
				FindDnsInAccount("test.com")
			}
			res := mylog.CaptureLogInfo(a)
			convey.So(res, convey.ShouldNotContainSubstring, "not")
		})
		convey.Convey("isDnsInAccount return false", func() {
			i.setDnsInAccountFlag(false)
			var a = func() {
				FindDnsInAccount("test.com")
			}
			res := mylog.CaptureLogInfo(a)
			convey.So(res, convey.ShouldContainSubstring, "not")
		})
	})
}
