package dns

import (
	"alitool/internal/pkg/test"
	"fmt"
	"testing"
)

func init() {
	test.GetEnv()
}

func Test_DescribeDomainRecordsViaA(t *testing.T) {
	//test.GetEnv()
	//viper.AutomaticEnv() // read in environment variables that match
	//viper.SetEnvPrefix("ALITOOL")
	//fmt.Println(viper.GetString("accessKeyId"))
	c := initDnsClient()
	fmt.Println(c.ListDomains())
}

func Test_DescribeDomainRecordsViaCNAME(t *testing.T) {
	c := initDnsClient()
	fmt.Printf("%#v", c.DescribeDomainRecordsViaCNAME("zhibosha.com"))
}
