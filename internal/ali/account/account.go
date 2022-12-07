package account

import (
	"github.com/spf13/viper"
)

// ali account map
// alias:{ak:xxx, sk:xxx}
var accountMap map[string]map[string]string

// var accounts []AliAccount
var accounts Accounts

type Accounts struct {
	AliAccount []aliAccount `mapstructure:"ali_account"`
}

type aliAccount struct {
	Alias           string `mapstructure:"alias"`
	AccessKeyId     string `mapstructure:"accessKeyId"`
	AccessKeySecret string `mapstructure:"secretKeyId"`
}

func GetAccountMap() map[string]map[string]string {
	return accountMap
}

// GetAccessKeyId return AccessKeyId
func (a *aliAccount) GetAccessKeyId() string {
	return a.AccessKeyId
}

// GetAccessKeySecret return AccessKeySecret
func (a *aliAccount) GetAccessKeySecret() string {
	return a.AccessKeySecret
}

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

func InitAccount() {
	accountUnmarshal()
	accountToMap()
}
