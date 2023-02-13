package dcdn

import . "alitool/internal/pkg/mylog"

func ListAllHttpsDomainsByAccount(d IDcdnClient) {
	res, keys := d.listAllHttpsDomains()
	for _, v := range keys {
		LoggerNoT.Printf("%s\n\t%+v\n", v, res[v])
	}
}

func ListAllHttpsDomainsInAllAccount() {
	initAllDcdnClient()
	allDcdnClient := getDcdnClients()
	for _, v := range allDcdnClient {
		LoggerNoT.Printf("\n\nAli account name: %s\n\n", v.getDcdnClientName())
		ListAllHttpsDomainsByAccount(v)
	}
}
