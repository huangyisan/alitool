package alego

import (
	"crypto"
	"encoding/json"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
	"github.com/sirupsen/logrus"
	"net/url"
	"os"
	"path/filepath"
)

// AccountsStorage A storage for account data.
//
// rootPath:
//
//	./.alego/accounts/
//	     │      └── root accounts directory
//	     └── "path" option
//
// rootUserPath:
//
//	./.alego/accounts/localhost_14000/hubert@hubert.com/
//	     │      │             │             └── userID ("email" option)
//	     │      │             └── CA server ("server" option)
//	     │      └── root accounts directory
//	     └── "path" option
//
// keysPath:
//
//	./.alego/accounts/localhost_14000/hubert@hubert.com/keys/
//	     │      │             │             │           └── root keys directory
//	     │      │             │             └── userID ("email" option)
//	     │      │             └── CA server ("server" option)
//	     │      └── root accounts directory
//	     └── "path" option
//
// accountFilePath:
//
//	./.alego/accounts/localhost_14000/hubert@hubert.com/account.json
//	     │      │             │             │             └── account file
//	     │      │             │             └── userID ("email" option)
//	     │      │             └── CA server ("server" option)
//	     │      └── root accounts directory
//	     └── "path" option

const (
	baseFolderName             = ".alego"
	baseAccountsRootFolderName = "accounts"
	baseKeysFolderName         = "keys"
	accountFileName            = "account.json"

	filePerm os.FileMode = 0o600
)

type AccountsStorage struct {
	userID          string
	rootPath        string
	rootUserPath    string
	keysPath        string
	accountFilePath string
}

func NewAccountsStorage(u IUser, acmeUrl string) *AccountsStorage {
	userID := u.GetEmail()
	rootPath := filepath.Join(os.Getenv("HOME"), baseFolderName, baseAccountsRootFolderName)
	urlPath, err := url.Parse(acmeUrl)
	if err != nil {
		logrus.Fatal(err)
	}
	rootUserPath := filepath.Join(rootPath, urlPath.Host, userID)
	accountFilePath := filepath.Join(rootUserPath, accountFileName)
	keysPath := filepath.Join(rootUserPath, baseKeysFolderName)

	return &AccountsStorage{
		userID:          userID,
		rootPath:        rootPath,
		rootUserPath:    rootUserPath,
		keysPath:        keysPath,
		accountFilePath: accountFilePath,
	}
}

func (s *AccountsStorage) ExistsAccountFilePath() bool {
	accountFile := filepath.Join(s.rootUserPath, accountFileName)
	if _, err := os.Stat(accountFile); os.IsNotExist(err) {
		return false
	} else if err != nil {
		logrus.Fatal(err)
	}
	return true
}

func (s *AccountsStorage) GetRootPath() string {
	return s.rootPath
}

func (s *AccountsStorage) GetRootUserPath() string {
	return s.rootUserPath
}

func (s *AccountsStorage) GetUserID() string {
	return s.userID
}

func createNonExistingFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0o700)
	} else if err != nil {
		return err
	}
	return nil
}

func (s *AccountsStorage) Save(account *AcmeUser) error {
	jsonBytes, err := json.MarshalIndent(account, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(s.accountFilePath, jsonBytes, filePerm)
}

func (s *AccountsStorage) LoadAccount() *AcmeUser {
	fileBytes, err := os.ReadFile(s.accountFilePath)
	if err != nil {
		logrus.Fatalf("Could not load file for account %s: %v", s.userID, err)
	}

	var account AcmeUser
	err = json.Unmarshal(fileBytes, &account)
	if err != nil {
		logrus.Fatalf("Could not parse file for account %s: %v", s.userID, err)
	}

	account.key = s.GetPrivateKey()
	//account.Email = s.userID

	if account.Registration == nil || account.Registration.Body.Status == "" {
		reg, err := tryRecoverRegistration(account.key)
		if err != nil {
			logrus.Fatalf("Could not load account for %s. Registration is nil: %#v", s.userID, err)
		}

		account.Registration = reg
		err = s.Save(&account)

		if err != nil {
			logrus.Fatalf("Could not save account for %s. Registration is nil: %#v", s.userID, err)
		}
	}
	return &account
}

func (s *AccountsStorage) RemoveRootUserPath() {
	os.RemoveAll(s.GetRootUserPath())
}

func tryRecoverRegistration(privateKey crypto.PrivateKey) (*registration.Resource, error) {
	// couldn't load account but got a key. Try to look the account up.
	config := lego.NewConfig(&AcmeUser{key: privateKey})
	config.CADirURL = "https://acme-staging-v02.api.letsencrypt.org/directory"
	config.UserAgent = "acm_go/0.0.1"

	client, err := lego.NewClient(config)
	if err != nil {
		return nil, err
	}

	reg, err := client.Registration.ResolveAccountByKey()
	if err != nil {
		return nil, err
	}
	return reg, nil
}
