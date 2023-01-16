package alego

import (
	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/lego"
	"github.com/sirupsen/logrus"
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

func CloudFlareVerification(client *lego.Client, authEmail, authToken string, domains ...string) {
	cloudFlareVerification(client, authEmail, authToken, domains...)
}
