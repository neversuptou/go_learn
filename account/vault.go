package account

import (
	"encoding/json"
	"github.com/fatih/color"
	"learnGO/utils"
	"time"
)

type ByteReader interface {
	Read() ([]byte, error)
}
type ByteWriter interface {
	Write(content []byte)
}

type Db interface {
	ByteWriter
	ByteReader
}
type Vault struct {
	Accounts  []AccountStruct `json:"accounts"`
	UpdatedAt time.Time       `json:"updatedAt"`
}
type VaultWithDB struct {
	Vault
	db Db
}

func NewVault(db Db) *VaultWithDB {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []AccountStruct{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		utils.PrintError("Error unmarshaling JSON")
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []AccountStruct{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	return &VaultWithDB{
		Vault: vault,
		db:    db,
	}
}

func (vault *VaultWithDB) addAccount(acc AccountStruct) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
	vault.save()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *VaultWithDB) FindAccounts(str string, checker func(AccountStruct, string) bool) []AccountStruct {
	var foundAccounts []AccountStruct
	for _, v := range vault.Accounts {
		//isMatched := strings.Contains(v.UrlString, urlString)
		isMatched := checker(v, str)
		if isMatched {
			foundAccounts = append(foundAccounts, v)
		}
	}
	return foundAccounts
}

func (vault *VaultWithDB) deleteAccountByURL(urlString string) bool {
	found := false
	for i, v := range vault.Accounts {
		if v.UrlString == urlString {
			vault.Accounts = append(vault.Accounts[:i], vault.Accounts[i+1:]...)
			vault.UpdatedAt = time.Now()
			found = true
			vault.save()
			break
		}
	}
	return found
}

func (vault *VaultWithDB) save() {
	file, err := vault.ToBytes()
	if err != nil {
		color.Red(err.Error())
	}
	vault.db.Write(file)
}
