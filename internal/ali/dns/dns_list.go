package dns

import "fmt"

func ListDnsByAccount(accountName string) {
	dnsClient := GetDnsClients()[accountName]
	hasRecordDomains := dnsClient.GetAllDomains()
	fmt.Printf("%s has dns record:\n", accountName)
	for record, _ := range hasRecordDomains {
		fmt.Println(record)
	}
}
