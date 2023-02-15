package account

import (
	"github.com/spf13/viper"
)

// ali account map
// alias:{ak:xxx, sk:xxx, subAccount: xxx}
var accountMap map[string]map[string]string

// var accounts []AliAccount
var accounts Accounts

type Accounts struct {
	AliAccount []AliAccount `mapstructure:"ali_account"`
}

type AliAccount struct {
	AccountName     string `mapstructure:"accountName"`
	AccessKeyId     string `mapstructure:"accessKeyId"`
	AccessKeySecret string `mapstructure:"secretKeyId"`
	SubAccount      string `mapstructure:"subAccount"`
}

func GetAccountMap() map[string]map[string]string {
	return accountMap
}

func (a *AliAccount) GetAccessKeyId() string {
	return a.AccessKeyId
}

func (a *AliAccount) GetAccessKeySecret() string {
	return a.AccessKeySecret
}

func accountUnmarshal() {
	viper.Unmarshal(&accounts)
}

func accountToMap() {
	accountMap = make(map[string]map[string]string)
	for _, v := range accounts.AliAccount {
		if accountMap[v.AccountName] == nil {
			accountMap[v.AccountName] = make(map[string]string)
		}
		accountMap[v.AccountName]["accessKeyId"] = v.AccessKeyId
		accountMap[v.AccountName]["accessKeySecret"] = v.AccessKeySecret
		accountMap[v.AccountName]["subAccount"] = v.SubAccount
	}
}

func InitAccount() {
	accountUnmarshal()
	accountToMap()
}
