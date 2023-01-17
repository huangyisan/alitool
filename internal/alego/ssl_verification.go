package alego

import (
	"fmt"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/providers/dns/cloudflare"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/url"
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
	certDomain := certificates.Domain
	certificateStorePath := filepath.Join(os.Getenv("HOME"), baseFolderName, baseAccountsRootFolderName, acmeApi, viper.GetString("acme.account.name"), certDomain)
	EnsureCertificatePath(certificateStorePath)
	err = os.WriteFile(fmt.Sprintf("%s/%s.crt", certificateStorePath, certDomain), certificates.Certificate, 0755)
	if err != nil {
		logrus.Fatal(err)
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s.key", certificateStorePath, certDomain), certificates.PrivateKey, 0755)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("store %s in %s", certDomain, certificateStorePath)
}

func getAcmeHost(client *lego.Client) string {
	regInfo, err := client.Registration.QueryRegistration()
	if err != nil {
		logrus.Fatal(err)
	}
	acmeUrl := regInfo.URI

	urlPath, err := url.Parse(acmeUrl)
	if err != nil {
		logrus.Fatal(err)
	}
	return urlPath.Host
}

func EnsureCertificatePath(certificateStorePath string) {
	if _, err := os.Stat(certificateStorePath); os.IsNotExist(err) {
		logrus.Infof("create certificate store path: %s\n", certificateStorePath)
		err := os.Mkdir(certificateStorePath, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}

// cloudFlareVerification use cloudflare dns verification
func cloudFlareVerification(client *lego.Client, authEmail, authToken string, domains ...string) {
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

	acmeHost := getAcmeHost(client)

	writeCertificateToFile(certificates, acmeHost)
}
