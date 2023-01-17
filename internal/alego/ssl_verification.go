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

// certificates storage
// rootUserPath:
//
//	./.alego/accounts/localhost_14000/hubert@hubert.com/
//	     │      │             │             └── userID ("email" option)
//	     │      │             └── CA server ("server" option)
//	     │      └── root accounts directory
//	     └── "path" option
//
// certificates path:
//
// ./.alego/accounts/localhost_14000/hubert@hubert.com/certificates/

const (
	cfDNSTtl                  = 120
	pollingInterval           = 2 * time.Second
	propagationTimeout        = 1 * time.Minute
	certificatesPath   string = "certificates"
)

func writeCertificateToFile(certificates *certificate.Resource, acmeApi string) {
	var err error
	certDomain := certificates.Domain
	certificateStorePath := filepath.Join(os.Getenv("HOME"), baseFolderName, baseAccountsRootFolderName, acmeApi, viper.GetString("acme.account.name"), certificatesPath, certDomain)
	ensureCertificatePath(certificateStorePath)
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

	return getUrlHost(acmeUrl)

}

func getUrlHost(urlPath string) string {
	u, err := url.Parse(urlPath)
	if err != nil {
		logrus.Fatal(err)
	}
	return u.Host
}

func ensureCertificatePath(certificateStorePath string) {
	if _, err := os.Stat(certificateStorePath); os.IsNotExist(err) {
		logrus.Infof("create certificate store path: %s\n", certificateStorePath)
		err := os.MkdirAll(certificateStorePath, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}

func listCertificates() []os.DirEntry {

	certificateStorePath := filepath.Join(os.Getenv("HOME"), baseFolderName, baseAccountsRootFolderName, getUrlHost(viper.GetString("acme.api_url.prod")), viper.GetString("acme.account.name"), certificatesPath)
	files, err := os.ReadDir(certificateStorePath)
	if err != nil {
		logrus.Fatal(err)
	}

	if len(files) > 0 {
		return files
	}
	return nil
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
