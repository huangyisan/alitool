package account

import (
	"fmt"
	"github.com/spf13/viper"
)

// ali account map
// alias:{ak:xxx, sk:xxx}
var accountMap map[string]map[string]string

// var accounts []aliAccount
var accounts Accounts

type Accounts struct {
	AliAccount []aliAccount `mapstructure:"ali_account"`
}

type aliAccount struct {
	Alias           string `mapstructure:"alias"`
	AccessKeyId     string `mapstructure:"accessKeyId"`
	AccessKeySecret string `mapstructure:"secretKeyId"`
}

//func GetAllAliAccounts() Accounts {
//	return accounts
//}

func accountUnmarshal() {
	viper.Unmarshal(&accounts)
}

func accountToMap() {
	accountMap = make(map[string]map[string]string)
	for _, v := range accounts.AliAccount {
		if accountMap[v.Alias] == nil {
			accountMap[v.Alias] = make(map[string]string)
		}
		accountMap[v.Alias]["accessKeyId"] = v.AccessKeyId
		accountMap[v.Alias]["accessKeySecret"] = v.AccessKeySecret
	}
}

func getAccountMap() map[string]map[string]string {
	return accountMap
}

func GetAccount(accountName string) (*aliAccount, bool) {
	v, ok := accountMap[accountName]
	if ok {
		return &aliAccount{
			Alias:           accountName,
			AccessKeyId:     v["AccessKeyId"],
			AccessKeySecret: v["AccessKeySecret"],
		}, true
	}
	fmt.Printf("cannot find %q account\n", accountName)
	DoListAccount()
	return nil, false
}

func InitAccount() {
	accountUnmarshal()
	accountToMap()
}