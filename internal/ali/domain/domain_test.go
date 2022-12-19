package domain

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/strategy"
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
	"github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func Test_newDomainClient(t *testing.T) {
	convey.Convey("Mock op.NewClient return nil", t, func() {
		var op *strategy.Operator
		patches := gomonkey.ApplyMethod(reflect.TypeOf(op), "NewClient", func(_ *strategy.Operator, _, _, _ string) (interface{}, error) {
			return "123", fmt.Errorf("Mock op.NewClient return nil\n")
		})
		defer patches.Reset()
		convey.Convey("Give params to newDnsClient", func() {
			res := newDomainClient("testName", "cn-shanghai", "ak", "sk")
			convey.So(res, convey.ShouldEqual, nil)
		})
	})
	convey.Convey("Mock op.NewClient return alidns.Client", t, func() {
		var op *strategy.Operator
		mockAccountName := "testName"
		mockRegion := "cn-shanghai"
		mockAK := "ak"
		mockSK := "sk"
		dc, _ := domain.NewClientWithAccessKey(mockRegion, mockAK, mockSK)

		patches := gomonkey.ApplyMethod(reflect.TypeOf(op), "NewClient", func(_ *strategy.Operator, _, _, _ string) (interface{}, error) {
			return dc, nil
		})
		defer patches.Reset()
		convey.Convey("Give params to newDnsClient", func() {
			res := newDomainClient("testName", "cn-shanghai", "ak", "sk")
			convey.So(res, convey.ShouldResemble, &DomainClient{
				AccountName: mockAccountName,
				RegionId:    mockRegion,
				I:           dc,
			})
		})
	})
	convey.Convey("Mock op.NewClient return not alidns.Client", t, func() {
		var op *strategy.Operator
		//mockAccountName := "testName"
		mockRegion := "cn-shanghai"
		mockAK := "ak"
		mockSK := "sk"
		// use alidns New func instead of dns
		dc, _ := alidns.NewClientWithAccessKey(mockRegion, mockAK, mockSK)

		patches := gomonkey.ApplyMethod(reflect.TypeOf(op), "NewClient", func(_ *strategy.Operator, _, _, _ string) (interface{}, error) {
			return dc, nil
		})
		defer patches.Reset()
		convey.Convey("Give params to newDnsClient", func() {
			res := newDomainClient("testName", "cn-shanghai", "ak", "sk")
			convey.So(res, convey.ShouldEqual, nil)
		})
	})
}

func TestInitDomainClient(t *testing.T) {
	wantAccountName := "account_01_patched"
	convey.Convey("Patched account.GetAccount func true", t, func() {
		patches := gomonkey.ApplyFunc(account.GetAccount, func(accountName string) (*account.AliAccount, bool) {
			return &account.AliAccount{
				AccountName:     wantAccountName,
				AccessKeyId:     "abc",
				AccessKeySecret: "def",
			}, true
		})
		defer patches.Reset()

		convey.Convey("Give accountName,", func() {
			dnsClient := InitDomainClient(wantAccountName, "cn-shanghai")
			wantClient := newDomainClient(wantAccountName, "cn-shanghai", "abc", "def")
			convey.So(dnsClient, convey.ShouldResemble, wantClient)
		})
	})

	convey.Convey("Patched account.GetAccount func false", t, func() {
		patches := gomonkey.ApplyFunc(account.GetAccount, func(accountName string) (*account.AliAccount, bool) {
			return &account.AliAccount{
				AccountName:     wantAccountName,
				AccessKeyId:     "abc",
				AccessKeySecret: "def",
			}, false
		})
		defer patches.Reset()

		convey.Convey("Give accountName,", func() {
			domainClient := InitDomainClient(wantAccountName, "cn-shanghai")
			convey.So(domainClient, convey.ShouldEqual, nil)
		})
	})
}

func Test_initAllDomainClient(t *testing.T) {
	convey.Convey("Mock account.GetAccountMap", t, func() {

		patches := gomonkey.ApplyFunc(account.GetAccountMap, func() map[string]map[string]string {
			return map[string]map[string]string{
				"account01": {
					"accessKeyId":     "accessKeyId01",
					"accessKeySecret": "accessKeySecret01",
				},
				"account02": {
					"accessKeyId":     "accessKeyId02",
					"accessKeySecret": "accessKeySecret02",
				}}
		})
		defer patches.Reset()

		convey.Convey("Mock account.AliAccount", func() {
			patches := gomonkey.ApplyFunc(account.GetAccount, func(string) (*account.AliAccount, bool) {
				return &account.AliAccount{
					AccountName:     "account01",
					AccessKeyId:     "accessKeyId01",
					AccessKeySecret: "accessKeySecret01",
				}, true
			})
			defer patches.Reset()

			convey.Convey("initAllDnsClients", func() {
				initAllDomainClient()
				want := []IDomainClient{
					InitDomainClient("account01", "cn-shanghai"),
					InitDomainClient("account02", "cn-shanghai"),
				}
				convey.So(domainClients, convey.ShouldResemble, want)
				// reset global var domainClients
				domainClients = make([]IDomainClient, 0)
			})
		})
	})
}

func Test_getDomainClients(t *testing.T) {
	tests := []struct {
		name string
		want []IDomainClient
	}{
		{
			name: "base case01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDomainClients(); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("getDnsClients() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAccountName(t *testing.T) {
	type fields struct {
		AccountName string
		RegionId    string
		I           iDomainClient
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "base case01",
			fields: fields{
				AccountName: "account01",
			},
			want: "account01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DomainClient{
				AccountName: tt.fields.AccountName,
				RegionId:    tt.fields.RegionId,
				I:           tt.fields.I,
			}
			if got := d.getAccountName(); got != tt.want {
				t.Errorf("getAccountName() = %v, want %v", got, tt.want)
			}
		})
	}
}
