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
var mockDnsRecord = "test.com"

type mockIDomainClient struct {
	domainInAccountFlag bool
}

func (m *mockIDomainClient) getAccountName() string {
	return mockAccountName
}

func (m *mockIDomainClient) listRegisteredDomainByAccount() recordRegisterDomains {
	//TODO implement me
	panic("implement me")
}

func (m *mockIDomainClient) isDomainInAccount(s string) bool {
	//TODO implement me
	return m.domainInAccountFlag
}

func (m *mockIDomainClient) getExpireDomains(i int) map[string]int {
	//TODO implement me
	panic("implement me")
}

func (m *mockIDomainClient) findExpireDomainRefAccount(s string) (string, int) {
	//TODO implement me
	panic("implement me")
}

func (m *mockIDomainClient) findExpireDomainsByAccount(client IDomainClient, i int) map[string]int {
	//TODO implement me
	panic("implement me")
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
