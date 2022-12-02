package account

func GetAccount(accountName string) (*aliAccount, bool) {
	v, ok := accountMap[accountName]
	if ok {
		return &aliAccount{
			Alias:           accountName,
			AccessKeyId:     v["accessKeyId"],
			AccessKeySecret: v["accessKeySecret"],
		}, true
	}
	return nil, false
}

func getAccountMap() map[string]map[string]string {
	return accountMap
}
