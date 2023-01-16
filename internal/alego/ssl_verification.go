package alego

import (
	"fmt"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/providers/dns/cloudflare"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"time"
)

const (
	cfDNSTtl           = 120
	pollingInterval    = 2 * time.Second
	propagationTimeout = 1 * time.Minute
)

func writeCertificateToFile(certificates *certificate.Resource, acmeApi string) {
	var err error
	certificateStorePath := filepath.Join(os.Getenv("HOME"), baseFolderName, baseAccountsRootFolderName, acmeApi, viper.GetString("acme.account.name"))
	err = os.WriteFile(fmt.Sprintf("%s/%s.crt", certificateStorePath, certificates.Domain), certificates.Certificate, 0755)
	if err != nil {
		logrus.Fatal(err)
	}
	err = os.WriteFile(fmt.Sprintf("%s/%s.key", certificateStorePath, certificates.Domain), certificates.PrivateKey, 0755)
	if err != nil {
		logrus.Fatal(err)
	}
}

// cloudFlareVerification use cloudflare dns verification
func cloudFlareVerification(client *lego.Client, authEmail, authToken string, acmeApi string, domains ...string) {
	var cloudFlareConfig cloudflare.Config
	cloudFlareConfig = cloudflare.Config{
		AuthEmail:          authEmail,
		AuthToken:          authToken,
		TTL:                cfDNSTtl,
		PollingInterval:    pollingInterval,
		PropagationTimeout: propagationTimeout,
	}

	cf, err := cloudflare.NewDNSProviderConfig(&cloudFlareConfig)
	if err != nil {
		logrus.Fatal(err)
	}

	err = client.Challenge.SetDNS01Provider(cf, dns01.AddDNSTimeout(2*time.Minute))
	if err != nil {
		logrus.Fatal(err)
	}

	request := certificate.ObtainRequest{
		Domains: domains,
		Bundle:  true,
	}

	certificates, err := client.Certificate.Obtain(request)
	if err != nil {
		logrus.Fatal(err)
	}
	writeCertificateToFile(certificates, acmeApi)
}
