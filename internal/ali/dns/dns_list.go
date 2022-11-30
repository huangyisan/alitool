package dns

import "fmt"

func ListDnsByAccount(accountName string) {
	dnsClient := GetDnsClients()
	fmt.Printf("tmp: %#v", dnsClient)
	//hasRecordDomains := dnsClient.GetAllDomains()
	//fmt.Printf("%s has dns record:\n", accountName)
	//for record, _ := range hasRecordDomains {
	//	fmt.Println(record)
	//}
}
