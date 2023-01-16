package alego

import (
	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
	"github.com/sirupsen/logrus"
)

func createUser(email string) *AcmeUser {
	return &AcmeUser{
		Email: email,
	}
}

func registrationAcmeAccount(email, CADirURL string) {
	user := createUser(email)
	config := lego.NewConfig(user)
	config.CADirURL = CADirURL
	config.Certificate.KeyType = certcrypto.RSA2048
	storage := NewAccountsStorage(user, config.CADirURL)
	privateKey := storage.GetPrivateKey()
	user.key = privateKey
	client, err := lego.NewClient(config)
	if err != nil {
		panic(err)
	}
	//register
	reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		// if error, then remove account root path
		storage.RemoveRootUserPath()
		logrus.Fatal(err)
	}
	// write registration info
	user.Registration = reg
	err = storage.Save(user)
	if err != nil {
		logrus.Fatal(err)
	}
}

func loadAcmeAccount(email, CADirURL string) *AcmeUser {
	user := createUser(email)
	storage := NewAccountsStorage(user, CADirURL)
	return storage.LoadAccount()
}
