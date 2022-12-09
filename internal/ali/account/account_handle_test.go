package account

import (
	"github.com/agiledragon/gomonkey/v2"
	"reflect"
	"testing"
)

func TestGetAccount(t *testing.T) {
	type args struct {
		accountName string
	}
	tests := []struct {
		name  string
		args  args
		want  *AliAccount
		want1 bool
	}{
		{"base case",
			struct{ accountName string }{accountName: "ali_account_01"}, &AliAccount{
				AccountName:     "ali_account_01",
				AccessKeyId:     "abc",
				AccessKeySecret: "def",
			}, true},
		{"wrong case",
			struct{ accountName string }{accountName: "ali_account_00"},
			nil,
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			patches := gomonkey.ApplyGlobalVar(&accountMap,
				map[string]map[string]string{"ali_account_01": {"accessKeyId": "abc", "accessKeySecret": "def"}})
			defer patches.Reset()
			got, got1 := GetAccount(tt.args.accountName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccount() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetAccount() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getAccountMap(t *testing.T) {
	tests := []struct {
		name string
		want map[string]map[string]string
	}{
		{
			name: "base case",
			want: map[string]map[string]string{"ali_account_01": {"accessKeyId": "abc", "accessKeySecret": "def"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			patches := gomonkey.ApplyGlobalVar(&accountMap,
				map[string]map[string]string{"ali_account_01": {"accessKeyId": "abc", "accessKeySecret": "def"}})
			defer patches.Reset()
			if got := getAccountMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAccountMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsExistAccount(t *testing.T) {
	type args struct {
		accountName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true case",
			args: args{
				accountName: "test_account_01_patched",
			},
			want: true,
		},
		{
			name: "false case",
			args: args{
				accountName: "test_account_02_patched",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			patches := gomonkey.ApplyGlobalVar(&accountMap, map[string]map[string]string{"test_account_01_patched": {}})
			defer patches.Reset()
			if got := IsExistAccount(tt.args.accountName); got != tt.want {
				t.Errorf("IsExistAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
