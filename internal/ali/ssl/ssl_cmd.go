package ssl

import (
	. "alitool/internal/pkg/mylog"
)

func GetExpireCertByAccount(i ISSLClient, expireDay int) {
	expireCerts := i.getExpireCertByAccount(expireDay)
	if expireCerts != nil {
		for k, v := range expireCerts {
			LoggerNoT.Printf("San %s found in %s account, expire in %d days, %s\n", k, i.getAccountName(), expireDay, v)
		}
	}
}

func GetExpireCertInAllAccounts(expireDay int) {
	initAllSSLClient()
	IFinanceClients := getSSLClients()
	for _, v := range IFinanceClients {
		GetExpireCertByAccount(v, expireDay)
	}
}

func GetExpireUploadCertByAccount(i ISSLClient, expireDay int) {
	expireCerts := i.getExpireUploadCertByAccount(expireDay)
	if expireCerts != nil {
		for k, v := range expireCerts {
			LoggerNoT.Printf("San %s found in %s account, expire in %d days, %s\n", k, i.getAccountName(), expireDay, v)
		}
	}
}

func GetExpireUploadCertInAllAccounts(expireDay int) {
	initAllSSLClient()
	IFinanceClients := getSSLClients()
	for _, v := range IFinanceClients {
		GetExpireUploadCertByAccount(v, expireDay)
	}
}
