package alego

import (
	. "alitool/internal/pkg/mylog"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/lego"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func NewAcmeClient(email string, CADirURL string) *lego.Client {
	return func(email string, CADirURL string) *lego.Client {
		acmeUser := loadAcmeAccount(email, CADirURL)
		config := lego.NewConfig(acmeUser)
		config.CADirURL = CADirURL
		config.Certificate.KeyType = certcrypto.RSA2048
		client, err := lego.NewClient(config)
		if err != nil {
			logrus.Fatal(err)
		}
		return client
	}(email, CADirURL)
}

func ListCertificates() {
	if files := listCertificates(); files != nil {
		LoggerNoT.Printf("List certificates:\n")
		for _, file := range files {
			LoggerNoT.Printf("%v\n", file.Name())
		}
	}
}

func ListCertificatesDetail() {
	certificateStorePath := filepath.Join(os.Getenv("HOME"), baseFolderName, baseAccountsRootFolderName, getUrlHost(viper.GetString("acme.api_url.prod")), viper.GetString("acme.account.name"), certificatesPath)
	if files := listCertificates(); files != nil {
		LoggerNoT.Printf("List certificates:\n")
		for _, file := range files {
			LoggerNoT.Printf("%v\n", file.Name())
			bs, err := os.ReadFile(filepath.Join(certificateStorePath, file.Name(), fmt.Sprintf("%s.crt", file.Name())))
			if err != nil {
				logrus.Fatal(err)
			}
			block, _ := pem.Decode(bs)
			if block == nil {
				logrus.Fatal("failed to parse PEM block containing the public key ->", fmt.Sprintf("%s.crt", file.Name()))
			}
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				logrus.Fatal(err)
			}
			LoggerNoT.Printf("\tIssuer: %+v\n", cert.Issuer)
			LoggerNoT.Printf("\tDNS names: %+v\n", cert.DNSNames)
			LoggerNoT.Printf("\tNot Before: %+v\n", cert.NotBefore)
			LoggerNoT.Printf("\tNot After: %+v\n", cert.NotAfter)
		}
	}
}

func CloudFlareVerification(client *lego.Client, authEmail, authToken string, domains ...string) {
	cloudFlareVerification(client, authEmail, authToken, domains...)
}
