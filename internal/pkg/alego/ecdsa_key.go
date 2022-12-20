package alego

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	privateKeyName string = "letsencrypt.key"
	publicKeyName  string = "letsencrypt.pub"
)

var keyPairPath string = fmt.Sprintf("%s/.ssh", os.Getenv("HOME"))

type ecdsaKey struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func genECDSAKey() *ecdsaKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	return &ecdsaKey{
		privateKey: privateKey,
		publicKey:  &privateKey.PublicKey,
	}
}

func makeECDSAKey() *ecdsaKey {
	if isKeyPairExist() {
		logrus.Infof("key pair in %s, read key pair to ecdsaKey...", keyPairPath)
		privateKey, publickKey := readFromFile()
		priv, pub := decodeKeyX509(privateKey, publickKey)
		return &ecdsaKey{
			privateKey: priv,
			publicKey:  pub,
		}
	}
	logrus.Infof("no key pair in %s, gen key pair...", keyPairPath)
	key := genECDSAKey()
	key.writeToFile()
	return key
}

func encodeKeyX509(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncoded), string(pemEncodedPub)
}

func decodeKeyX509(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return privateKey, publicKey
}

func (e *ecdsaKey) writeToFile() {
	encPriv, encPub := encodeKeyX509(e.privateKey, e.publicKey)
	encPrivByte := []byte(encPriv)
	encPubByte := []byte(encPub)
	os.WriteFile(fmt.Sprintf("%s/%s", keyPairPath, privateKeyName), encPrivByte, 400)
	os.WriteFile(fmt.Sprintf("%s/%s", keyPairPath, publicKeyName), encPubByte, 400)
}

func readFromFile() (privateKey, publicKey string) {
	privateKeyByte, err := os.ReadFile(fmt.Sprintf("%s/%s", keyPairPath, privateKeyName))
	if err != nil {
		logrus.Fatalln(err)
	}
	publicKeyByte, err := os.ReadFile(fmt.Sprintf("%s/%s", keyPairPath, publicKeyName))
	if err != nil {
		logrus.Fatalln(err)
	}
	privateKey = string(privateKeyByte)
	publicKey = string(publicKeyByte)
	return
}

func isKeyPairExist() bool {
	priv, err := os.Stat(fmt.Sprintf("%s/%s", keyPairPath, privateKeyName))
	if err != nil {
		return false
	}
	pub, err := os.Stat(fmt.Sprintf("%s/%s", keyPairPath, publicKeyName))
	if err != nil {
		return false
	}
	if !priv.IsDir() && !pub.IsDir() {
		return true
	}
	return false
}
