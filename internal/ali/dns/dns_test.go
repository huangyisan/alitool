package dns

import (
	"alitool/internal/ali/account"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func Test_newDnsClient(t *testing.T) {
	type args struct {
		regionId        string
		accessKeyId     string
		accessKeySecret string
		accountName     string
	}
	tests := []struct {
		name     string
		mockFunc func(accountName, regionId, accessKeyId, accessKeySecret string) IDNSClient
		args     args
		want     IDNSClient
	}{
		{
			name: "base case01", args: args{
				regionId:        "cn-shanghai",
				accessKeyId:     "abc",
				accessKeySecret: "def",
			},
			mockFunc: func(accountName, regionId, accessKeyId, accessKeySecret string) IDNSClient {
				return newDnsClient(accountName, regionId, accessKeyId, accessKeySecret)
			},
			want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = tt.mockFunc(tt.args.accountName, tt.args.regionId, tt.args.accessKeyId, tt.args.accessKeySecret)
			if got := newDnsClient(tt.args.accountName, tt.args.regionId, tt.args.accessKeyId, tt.args.accessKeySecret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDnsClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitDnsClient(t *testing.T) {
	wantAccountName := "account_01_patched"
	convey.Convey("Patched account.GetAccount func", t, func() {
		patches := gomonkey.ApplyFunc(account.GetAccount, func(accountName string) (*account.AliAccount, bool) {
			return &account.AliAccount{
				AccountName:     wantAccountName,
				AccessKeyId:     "abc",
				AccessKeySecret: "def",
			}, true
		})
		defer patches.Reset()

		convey.Convey("Give accountName,", func() {
			dnsClient := InitDnsClient(wantAccountName, "cn-shanghai")
			wantClient := newDnsClient(wantAccountName, "cn-shanghai", "abc", "def")
			convey.So(dnsClient, convey.ShouldResemble, wantClient)
		})
	})
}

func TestDnsClient_getAccountName(t *testing.T) {
	type fields struct {
		AccountName string
		RegionId    string
		I           iDnsClient
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
			d := &DnsClient{
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

func Test_getDnsClients(t *testing.T) {
	tests := []struct {
		name string
		want []IDNSClient
	}{
		{
			name: "base case01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDnsClients(); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("getDnsClients() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initAllDnsClients(t *testing.T) {
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
				initAllDnsClients()
				want := []IDNSClient{
					InitDnsClient("account01", "cn-shanghai"),
					InitDnsClient("account02", "cn-shanghai"),
				}
				convey.So(dnsClients, convey.ShouldResemble, want)
			})
		})
	})
}
