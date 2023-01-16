package alego

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type ecdsaKey struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func genECDSAKey() (*ecdsaKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	return &ecdsaKey{
		privateKey: privateKey,
		publicKey:  &privateKey.PublicKey,
	}, nil
}

func generateECDSAKey(keyPath string) (*ecdsaKey, error) {
	key, err := genECDSAKey()
	if err != nil {
		return nil, err
	}
	err = key.writeToFile(keyPath)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func encodeKeyX509(privateKey *ecdsa.PrivateKey) (string, error) {
	x509Encoded, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return "", err
	}
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	return string(pemEncoded), nil
}

func decodeKeyX509(pemEncoded string) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, err := x509.ParseECPrivateKey(x509Encoded)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func (e *ecdsaKey) writeToFile(accKeyPath string) error {
	encPriv, err := encodeKeyX509(e.privateKey)
	if err != nil {
		return err
	}
	encPrivByte := []byte(encPriv)
	err = os.WriteFile(accKeyPath, encPrivByte, 400)
	if err != nil {
		return err
	}
	return nil
}

func readFromFile(accKeyPath string) (privateKey string) {
	privateKeyByte, err := os.ReadFile(accKeyPath)
	if err != nil {
		logrus.Fatalln(err)
	}

	privateKey = string(privateKeyByte)
	return
}

func (s *AccountsStorage) createKeysFolder() {
	if err := createNonExistingFolder(s.keysPath); err != nil {
		logrus.Fatalf("Could not check/create directory for account %s: %v", s.userID, err)
	}
}

func (s *AccountsStorage) GetPrivateKey() *ecdsa.PrivateKey {
	accKeyPath := filepath.Join(s.keysPath, s.userID+".key")

	if _, err := os.Stat(accKeyPath); os.IsNotExist(err) {
		logrus.Printf("No key found for account %s. Generating a %s key.", s.userID, "ECDSA")
		s.createKeysFolder()

		keyPair, err := generateECDSAKey(accKeyPath)
		if err != nil {
			logrus.Fatalf("Could not generate RSA private account key for account %s: %v", s.userID, err)
		}

		logrus.Printf("Saved key to %s", accKeyPath)
		return keyPair.privateKey
	}

	privateKey := readFromFile(accKeyPath)
	priv, err := decodeKeyX509(privateKey)
	if err != nil {
		logrus.Fatalf("Could not load RSA private key from file %s: %v", accKeyPath, err)
	}
	return priv
}
