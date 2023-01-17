package alego

import (
	"crypto"
	"github.com/go-acme/lego/v4/registration"
	"github.com/sirupsen/logrus"
	"net/url"
)

// You'll need a user or account type that implements acme.User
type AcmeUser struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

func (u *AcmeUser) GetEmail() string {
	return u.Email
}
func (u AcmeUser) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *AcmeUser) GetPrivateKey() crypto.PrivateKey {
	return u.key
	//p, _ := genECDSAKey()
	//return p.privateKey
}

func (u *AcmeUser) GetRegistrationUrl() string {
	uri, err := url.Parse(u.Registration.URI)
	if err != nil {
		logrus.Fatal(err)
	}

	return uri.Host
}

type IUser interface {
	GetEmail() string
	GetRegistration() *registration.Resource
	GetPrivateKey() crypto.PrivateKey
	GetRegistrationUrl() string
}
