package domain

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/mylog"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func Test_listRegisteredDomainByAccount(t *testing.T) {
	var d *DomainClient
	convey.Convey("Mock getAllDnsDomains", t, func() {
		patches := gomonkey.ApplyPrivateMethod(reflect.TypeOf(d), "getAllRegisteredDomains", func(*DomainClient) recordRegisterDomains {
			return recordRegisterDomains{"test.com": struct{}{}}
		})
		defer patches.Reset()
		convey.Convey("listDnsByAccount", func() {
			hasRecordDomains := d.listRegisteredDomainByAccount()
			convey.So(hasRecordDomains, convey.ShouldResemble, recordRegisterDomains{"test.com": struct{}{}})
		})
	})
}

func Test_isDomainInAccount(t *testing.T) {
	var d *DomainClient
	trueCaseDomain := "test.com"
	falseCaseDomain := "false.com"
	convey.Convey("Mock listDnsByAccount", t, func() {
		patches := gomonkey.ApplyPrivateMethod(reflect.TypeOf(d), "listRegisteredDomainByAccount", func(_ *DomainClient) recordRegisterDomains {
			return recordRegisterDomains{"test.com": struct{}{}}
		})
		defer patches.Reset()

		convey.Convey("Give trueCaseDomain", func() {
			res := d.isDomainInAccount(trueCaseDomain)
			convey.So(res, convey.ShouldEqual, true)
		})

		convey.Convey("Give falseCaseDomain", func() {
			res := d.isDomainInAccount(falseCaseDomain)
			convey.So(res, convey.ShouldEqual, false)
		})
	})
}

var mockAccountName = "mockAccountName"
var mockDomainRecord = "test.com"
var mockExpireDay = 10
var mockExpireDomainsInfo = expireDomainsInfo{mockDomainRecord: mockExpireDay}
var mockRecordRegisterDomains = recordRegisterDomains{mockDomainRecord: struct{}{}}

type mockIDomainClient struct {
	domainInAccountFlag      bool
	recordRegisterDomainsLen int
	expireDomainsInfoLen     int
}

func (m *mockIDomainClient) getAccountName() string {
	return mockAccountName
}

func (m *mockIDomainClient) listRegisteredDomainByAccount() recordRegisterDomains {
	if m.recordRegisterDomainsLen == 0 {
		return recordRegisterDomains{}
	}
	return mockRecordRegisterDomains

}

func (m *mockIDomainClient) isDomainInAccount(s string) bool {
	return m.domainInAccountFlag
}

func (m *mockIDomainClient) getExpireDomains(i int) expireDomainsInfo {
	return map[string]int{"www.test.com": i}
}

func (m *mockIDomainClient) findExpireDomainRefAccount(s string) (string, int) {
	if s == mockDomainRecord {
		return mockDomainRecord, mockExpireDay
	}
	return mockDomainRecord, -1
}

func (m *mockIDomainClient) findExpireDomainsByAccount(i int) expireDomainsInfo {
	if m.expireDomainsInfoLen > 0 {
		return expireDomainsInfo{mockDomainRecord: mockExpireDay}
	}
	return expireDomainsInfo{}
}

func (m *mockIDomainClient) setRecordRegisterDomainsLen(i int) {
	m.recordRegisterDomainsLen = i
}

func (m *mockIDomainClient) setExpireDomainsInfoLen(i int) {
	m.expireDomainsInfoLen = i
}

func (m *mockIDomainClient) setDomainInAccountFlag(flag bool) {
	m.domainInAccountFlag = flag
}

func TestIsDomainInAccount(t *testing.T) {
	convey.Convey("Mock account.IsExistAccount return false", t, func() {
		patches := gomonkey.ApplyFunc(account.IsExistAccount, func(accountName string) bool {
			return false
		})
		defer patches.Reset()

		var i mockIDomainClient
		domainName := "test.com"
		var a = func() {
			func(i IDomainClient, domainName string) {
				IsDomainInAccount(i, domainName)
			}(&i, domainName)
		}
		res := mylog.CaptureLogInfo(a)
		convey.So(res, convey.ShouldContainSubstring, "not")

	})
	convey.Convey("Mock account.IsExistAccount return true", t, func() {
		patches := gomonkey.ApplyFunc(account.IsExistAccount, func(accountName string) bool {
			return true
		})
		defer patches.Reset()
		convey.Convey("Give invalid domainName", func() {
			var i mockIDomainClient
			domainName := "com"
			var a = func() {
				func(i IDomainClient, domainName string) {
					IsDomainInAccount(i, domainName)
				}(&i, domainName)
			}
			res := mylog.CaptureLogInfo(a)
			convey.So(res, convey.ShouldContainSubstring, "not")

		})

		convey.Convey("Give valid domainName", func() {
			convey.Convey("dns in account", func() {
				//var i mockIDNSClient
				domainName := "test.com"

				var a = func() {
					i := mockIDomainClient{}
					i.setDomainInAccountFlag(true)
					func(i IDomainClient, domainName string) {
						IsDomainInAccount(i, domainName)
					}(&i, domainName)
				}
				res := mylog.CaptureLogInfo(a)
				convey.So(res, convey.ShouldNotContainSubstring, "not")

			})

			convey.Convey("dns not in account", func() {
				var i mockIDomainClient
				domainName := "test.com"
				i.setDomainInAccountFlag(false)
				var a = func() {
					func(i IDomainClient, domainName string) {
						IsDomainInAccount(i, domainName)
					}(&i, domainName)
				}
				res := mylog.CaptureLogInfo(a)
				convey.So(res, convey.ShouldContainSubstring, "not")
			})
		})
	})
}

func TestFindDomainInAccount(t *testing.T) {
	convey.Convey("Mock getDomainClients []IDomainClient{} is empty", t, func() {
		patches := gomonkey.ApplyFunc(getDomainClients, func() []IDomainClient {
			return []IDomainClient{}
		})
		defer patches.Reset()
		var a = func() {
			FindDomainInAccount("test.com")
		}
		res := mylog.CaptureLogInfo(a)
		convey.So(res, convey.ShouldContainSubstring, "not")
	})

	convey.Convey("Mock getDnsClients []IDNSClient{} is not empty", t, func() {
		var i mockIDomainClient
		var c mockIDomainClient
		patches := gomonkey.ApplyFunc(getDomainClients, func() []IDomainClient {
			return []IDomainClient{&i, &c}
		})
		defer patches.Reset()
		convey.Convey("isDnsInAccount return true", func() {
			i.setDomainInAccountFlag(true)
			var a = func() {
				FindDomainInAccount("test.com")
			}
			res := mylog.CaptureLogInfo(a)
			convey.So(res, convey.ShouldNotContainSubstring, "not")
		})
		convey.Convey("isDnsInAccount return false", func() {
			i.setDomainInAccountFlag(false)
			var a = func() {
				FindDomainInAccount("test.com")
			}
			res := mylog.CaptureLogInfo(a)
			convey.So(res, convey.ShouldContainSubstring, "not")
		})
	})
}

func Test_findExpireDomainsByAccount(t *testing.T) {
	var d *DomainClient
	convey.Convey("Mock getExpireDomains", t, func() {
		patches := gomonkey.ApplyPrivateMethod(d, "getExpireDomains", func(_ *DomainClient, i int) expireDomainsInfo {
			return mockExpireDomainsInfo
		})
		defer patches.Reset()
		res := d.findExpireDomainsByAccount(mockExpireDay)
		convey.So(res, convey.ShouldResemble, mockExpireDomainsInfo)
	})
}

func Test_findExpireDomainRefAccount(t *testing.T) {
	d := &DomainClient{
		AccountName: mockAccountName,
	}
	mockDomain := "www.test.com"
	convey.Convey("Mock isDomainInAccount return false", t, func() {
		patches := gomonkey.ApplyPrivateMethod(d, "isDomainInAccount", func(_ *DomainClient, _ string) bool {
			return false
		})
		defer patches.Reset()
		accountName, expireDay := d.findExpireDomainRefAccount(mockDomain)
		convey.So(accountName, convey.ShouldEqual, mockAccountName)
		convey.So(expireDay, convey.ShouldEqual, -1)
	})

	convey.Convey("Mock isDomainInAccount return true", t, func() {
		patches := gomonkey.ApplyPrivateMethod(d, "isDomainInAccount", func(_ *DomainClient, _ string) bool {
			return true
		})
		defer patches.Reset()
		convey.Convey("Mock getDomainExpireCurrDiff return mockExpireDay", func() {
			patches := gomonkey.ApplyPrivateMethod(d, "getDomainExpireCurrDiff", func(_ *DomainClient, string, _ int) int {
				return mockExpireDay
			})
			defer patches.Reset()
			accountName, expireDay := d.findExpireDomainRefAccount(mockDomain)
			convey.So(accountName, convey.ShouldEqual, mockAccountName)
			convey.So(expireDay, convey.ShouldEqual, mockExpireDay)
		})

	})
}

func TestFindExpireDomainRefAccount(t *testing.T) {
	d := mockIDomainClient{}
	mockNotExistDomainRecord := "notexist.com"
	convey.Convey("accountName is \"\" ", t, func() {
		var a = func() {
			FindExpireDomainRefAccount(&d, mockNotExistDomainRecord)
		}
		res := mylog.CaptureLogInfo(a)
		convey.So(res, convey.ShouldContainSubstring, "no")
	})

	convey.Convey("accountName is \"\" ", t, func() {
		var a = func() {
			FindExpireDomainRefAccount(&d, mockDomainRecord)
		}
		res := mylog.CaptureLogInfo(a)
		convey.So(res, convey.ShouldNotContainSubstring, "no")
	})
}

func TestListRegisteredDomainByAccount(t *testing.T) {
	d := mockIDomainClient{}
	convey.Convey("Mock recordRegisterDomainsLen == 0 ", t, func() {
		d.setRecordRegisterDomainsLen(0)
		var a = func() {
			ListRegisteredDomainByAccount(&d)
		}
		res := mylog.CaptureLogInfo(a)
		convey.So(res, convey.ShouldContainSubstring, "Total count")
	})
	convey.Convey("Mock recordRegisterDomainsLen > 0 ", t, func() {
		d.setRecordRegisterDomainsLen(1)
		var a = func() {
			ListRegisteredDomainByAccount(&d)
		}
		res := mylog.CaptureLogInfo(a)
		convey.So(res, convey.ShouldContainSubstring, "exist")
	})
}

func TestFindExpireDomainsByAccount(t *testing.T) {
	d := mockIDomainClient{}
	convey.Convey("expireDomains len is 0", t, func() {
		d.setExpireDomainsInfoLen(0)
		var a = func() {
			FindExpireDomainsByAccount(&d, mockExpireDay)
		}
		res := mylog.CaptureLogInfo(a)
		convey.So(res, convey.ShouldContainSubstring, "no")
	})
	convey.Convey("expireDomains len i> 0", t, func() {
		d.setExpireDomainsInfoLen(1)
		var a = func() {
			FindExpireDomainsByAccount(&d, mockExpireDay)
		}
		res := mylog.CaptureLogInfo(a)
		convey.So(res, convey.ShouldNotContainSubstring, "no")
	})
}

func TestFindExpireDomainsInAllAccounts(t *testing.T) {
	var d mockIDomainClient
	convey.Convey("Mock initAllDomainClient", t, func() {
		patches := gomonkey.ApplyFunc(initAllDomainClient, func() {

		})
		defer patches.Reset()
		convey.Convey("expireDomains > 0", func() {
			convey.Convey("Mock domainClient", func() {
				patches := gomonkey.ApplyGlobalVar(&domainClients, []IDomainClient{&d})
				defer patches.Reset()
				d.setExpireDomainsInfoLen(1)
				var a = func() {
					FindExpireDomainsInAllAccounts(mockExpireDay)
				}
				res := mylog.CaptureLogInfo(a)
				convey.So(res, convey.ShouldNotContainSubstring, "no")
			})

		})
	})
}
