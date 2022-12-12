package dns

import (
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
