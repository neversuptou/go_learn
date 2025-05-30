package account

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"learnGO/files"
	"time"
)

type Vault struct {
	Accounts  []AccountStruct `json:"accounts"`
	UpdatedAt time.Time       `json:"updatedAt"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []AccountStruct{},
			UpdatedAt: time.Now(),
		}
	}

	var vault Vault

	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Error unmarshaling JSON", err)
		return &Vault{
			Accounts:  []AccountStruct{},
			UpdatedAt: time.Now(),
		}
	}
	return &vault
}

func (vault *Vault) addAccount(acc AccountStruct) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()

	data, err := vault.ToBytes()
	if err != nil {
		color.Red("\"Error marshaling JSON\"", err)
	}
	files.SaveFile(data, "data.json")
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func findAccountByURL(urlString string) {
	found := false
	for i, v := range accountArray {
		if v.UrlString == urlString {
			fmt.Printf("Found account:\nID:%d \nLogin: %s\nPassword: %s\nURL: %s\n",
				i+1, v.Login, v.Password, v.UrlString)
			found = true
		}
	}
	if !found {
		fmt.Println("Account is not found")
	}
}
