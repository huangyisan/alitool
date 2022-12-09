package dns

import (
	"alitool/internal/ali/account"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func setup() {
	//test.GetEnv()
	//account.InitAccount()
}

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
	var i IDNSClient
	convey.Convey("Patched account.GetAccount func", t, func() {

		var mockGetAccount = gomonkey.ApplyFunc(account.GetAccount, func(accountName string) (*account.AliAccount, bool) {
			return &account.AliAccount{
				AccountName:     "account_01_patched",
				AccessKeyId:     "abc",
				AccessKeySecret: "def",
			}, true
		})
		defer mockGetAccount.Reset()
		convey.Convey("Give accountName,", func() {
			dnsClient := InitDnsClient("account_01_patched", "cn-shanghai")
			convey.So(dnsClient, convey.ShouldEqual, i)
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
