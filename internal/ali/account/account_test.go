package account

import (
	"alitool/internal/pkg/test"
	"reflect"
	"testing"
)

func setup() {
	test.GetEnv()
	InitAccount()
}
func Test_GetAccount(t *testing.T) {
	//fmt.Printf("%#v", accounts)

	GetAccount("ali_account_01")
}

func TestGetAccount(t *testing.T) {
	setup()
	type args struct {
		accountName string
	}
	tests := []struct {
		name  string
		args  args
		want  *aliAccount
		want1 bool
	}{
		{"base case", struct{ accountName string }{accountName: "ali_account_01"}, &aliAccount{
			Alias:           "ali_account_01",
			AccessKeyId:     "abc",
			AccessKeySecret: "def",
		}, true},
		{"wrong case", struct{ accountName string }{accountName: "ali_account_00"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
