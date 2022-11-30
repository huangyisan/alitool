package dns

import (
	"alitool/internal/ali/account"
	"alitool/internal/pkg/test"
	"reflect"
	"testing"
)

func setup() {
	test.GetEnv()
	account.InitAccount()
	initAllDnsClient()
}

func Test_initDnsClient(t *testing.T) {
	// setup env
	setup()

	type args struct {
		accountName     string
		regionId        string
		accessKeyId     string
		accessKeySecret string
	}

	tests := []struct {
		name     string
		args     args
		mockFunc func(regionId, accessKeyId, accessKeySecret string) *DnsClient
		want     *DnsClient
	}{
		{
			name: "base case01",
			args: args{
				accountName:     "ali_account_01",
				regionId:        "cn-hangzhou",
				accessKeyId:     "abc",
				accessKeySecret: "def",
			},
			mockFunc: func(regionId, accessKeyId, accessKeySecret string) *DnsClient {
				return NewDnsClient(regionId, accessKeyId, accessKeySecret)
			},
			want: &DnsClient{},
		},
		{
			name: "base case02",
			args: args{
				accountName:     "ali_account_99",
				regionId:        "cn-shanghai",
				accessKeyId:     "abc",
				accessKeySecret: "def",
			},
			mockFunc: func(regionId, accessKeyId, accessKeySecret string) *DnsClient {
				return nil
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = tt.mockFunc(tt.args.regionId, tt.args.accessKeyId, tt.args.accessKeySecret)
			if got := initDnsClient(tt.args.accountName, tt.args.regionId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initDnsClient() = %#v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDnsClient(t *testing.T) {
	setup()
	type args struct {
		regionId        string
		accessKeyId     string
		accessKeySecret string
	}
	tests := []struct {
		name     string
		mockFunc func(regionId, accessKeyId, accessKeySecret string) *DnsClient
		args     args
		want     *DnsClient
	}{
		{
			name: "base case01", args: args{
				regionId:        "cn-shanghai",
				accessKeyId:     "abc",
				accessKeySecret: "def",
			},
			mockFunc: func(regionId, accessKeyId, accessKeySecret string) *DnsClient {
				return NewDnsClient(regionId, accessKeyId, accessKeySecret)
			},
			want: nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = tt.mockFunc(tt.args.regionId, tt.args.accessKeyId, tt.args.accessKeySecret)
			if got := NewDnsClient(tt.args.regionId, tt.args.accessKeyId, tt.args.accessKeySecret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDnsClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initAllDnsClient(t *testing.T) {
	setup()
	initAllDnsClient()
}
