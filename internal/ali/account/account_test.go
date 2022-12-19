package account

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_GetAccessKeyId(t *testing.T) {
	type fields struct {
		Alias           string
		AccessKeyId     string
		AccessKeySecret string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "base case 01",
			fields: fields{
				Alias:           "baseCase01",
				AccessKeyId:     "baseCase01AccessKeyId",
				AccessKeySecret: "baseCase01AccessKeySecret",
			},
			want: "baseCase01AccessKeyId",
		},
		{
			name: "base case 02",
			fields: fields{
				Alias:           "baseCase02",
				AccessKeyId:     "baseCase02AccessKeyId",
				AccessKeySecret: "baseCase02AccessKeySecret",
			},
			want: "baseCase02AccessKeyId",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AliAccount{
				AccountName:     tt.fields.Alias,
				AccessKeyId:     tt.fields.AccessKeyId,
				AccessKeySecret: tt.fields.AccessKeySecret,
			}
			if got := a.GetAccessKeyId(); got != tt.want {
				t.Errorf("GetAccessKeyId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetAccessKeySecret(t *testing.T) {
	type fields struct {
		Alias           string
		AccessKeyId     string
		AccessKeySecret string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "base case 01",
			fields: fields{
				Alias:           "baseCase01",
				AccessKeyId:     "baseCase01AccessKeyId",
				AccessKeySecret: "baseCase01AccessKeySecret",
			},
			want: "baseCase01AccessKeySecret",
		},
		{
			name: "base case 02",
			fields: fields{
				Alias:           "baseCase02",
				AccessKeyId:     "baseCase02AccessKeyId",
				AccessKeySecret: "baseCase02AccessKeySecret",
			},
			want: "baseCase02AccessKeySecret",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AliAccount{
				AccountName:     tt.fields.Alias,
				AccessKeyId:     tt.fields.AccessKeyId,
				AccessKeySecret: tt.fields.AccessKeySecret,
			}
			if got := a.GetAccessKeySecret(); got != tt.want {
				t.Errorf("GetAccessKeySecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAccountMap(t *testing.T) {
	convey.Convey("Mock accountMap", t, func() {
		patches := gomonkey.ApplyGlobalVar(&accountMap, map[string]map[string]string{"account01": {}})
		defer patches.Reset()
		convey.Convey("GetAccountMap", func() {
			res := GetAccountMap()
			convey.So(res, convey.ShouldResemble, accountMap)
		})
	})
}

func Test_accountToMap(t *testing.T) {
	wantAccount01 := "accountName01"
	wantAccessKey01 := "accessKeyId01"
	wantAccessKeySecret01 := "accessKeySecret01"

	wantAccount02 := "accountName02"
	wantAccessKey02 := "accessKeyId02"
	wantAccessKeySecret02 := "accessKeySecret02"

	convey.Convey("Mock accounts", t, func() {
		patches := gomonkey.ApplyGlobalVar(&accounts, Accounts{AliAccount: []AliAccount{
			{
				AccountName:     wantAccount01,
				AccessKeyId:     wantAccessKey01,
				AccessKeySecret: wantAccessKeySecret01,
			},
			{
				AccountName:     wantAccount02,
				AccessKeyId:     wantAccessKey02,
				AccessKeySecret: wantAccessKeySecret02,
			},
		}})
		want := map[string]map[string]string{
			wantAccount01: {
				"accessKeyId":     wantAccessKey01,
				"accessKeySecret": wantAccessKeySecret01,
			},
			wantAccount02: {
				"accessKeyId":     wantAccessKey02,
				"accessKeySecret": wantAccessKeySecret02,
			},
		}
		defer patches.Reset()
		convey.Convey("accountToMap", func() {
			accountToMap()
			convey.So(accountMap, convey.ShouldResemble, want)
		})
	})
}
