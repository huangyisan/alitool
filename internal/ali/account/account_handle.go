package account

// GetAccount will return a AliAccount
func GetAccount(accountName string) (*AliAccount, bool) {
	v, ok := accountMap[accountName]
	if ok {
		return &AliAccount{
			AccountName:     accountName,
			AccessKeyId:     v["accessKeyId"],
			AccessKeySecret: v["accessKeySecret"],
		}, true
	}
	return nil, false
}

// getAccountMap Will return accountMap
func getAccountMap() map[string]map[string]string {
	return accountMap
}

// IsExistAccount will return true if account in accountMap, else false
func IsExistAccount(accountName string) bool {
	_, ok := GetAccount(accountName)
	if ok {
		return true
	}
	return false
}
