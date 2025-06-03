package account

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"learnGO/utils"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

type AccountStruct struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	UrlString string    `json:"urlString"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var accountArray = []AccountStruct{}

func (acc *AccountStruct) generatePassword(n int) {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789@-_")
	arr := make([]rune, n)

	for i := range arr {
		arr[i] = letters[rand.Intn(len(letters))]
	}
	acc.Password = string(arr)
}

func NewAccount(login, password, urlString string) (*AccountStruct, error) {
	if login == "" {
		return nil, errors.New("LOGIN_INVALID")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("URL_INVALID")
	}

	newAcc := &AccountStruct{
		Login:     login,
		Password:  password,
		UrlString: urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if password == "" {
		newAcc.generatePassword(10)
	}
	accountArray = append(accountArray, *newAcc)
	return newAcc, nil
}

func FindAccountWithUrl(vault *VaultWithDB) {
	if len(vault.Accounts) == 0 {
		utils.PrintError("Accounts not created")
		return
	}
	urlString := utils.PromptData("Enter url of the account you want to find: ")
	foundAccounts := vault.FindAccounts(urlString, func(acc AccountStruct, str string) bool {
		return strings.Contains(acc.UrlString, str)
	})
	outPutAccounts(&foundAccounts, urlString)
}

func FindAccountWithLogin(vault *VaultWithDB) {
	if len(vault.Accounts) == 0 {
		utils.PrintError("Accounts not created")
		return
	}
	login := utils.PromptData("Enter login of the account you want to find: ")
	foundAccounts := vault.FindAccounts(login, func(acc AccountStruct, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outPutAccounts(&foundAccounts, login)
}

func outPutAccounts(foundAccounts *[]AccountStruct, str string) {
	if len(*foundAccounts) == 0 {
		utils.PrintError("Accounts not found with login: " + str)
		return
	}
	for _, v := range *foundAccounts {
		output(&v, "Account: ")
	}
}

func DeleteAccount(vault *VaultWithDB) {
	if len(vault.Accounts) == 0 {
		utils.PrintError("Accounts not created")
		return
	}
	urlString := utils.PromptData("Enter url of the account you want to delete: ")
	vault.UpdatedAt = time.Now()

	if !vault.deleteAccountByURL(urlString) {
		utils.PrintError("Account not found")
		return
	}
	color.Green("Account deleted successfully")
}

func CreateAccount(vault *VaultWithDB) {
	login := utils.PromptData("Enter your login: ")
	password := utils.PromptData("Enter your Password: ")
	urlString := utils.PromptData("Enter your URL: ")

	dataAccount, err := NewAccount(login, password, urlString)
	if err != nil {
		fmt.Println(err)
		return
	}

	vault.addAccount(*dataAccount)

	output(dataAccount, "Account saved successfully:")
}

func output(acc *AccountStruct, text string) {
	color.Cyan(text)
	color.Cyan("Login: " + acc.Login)
	color.Cyan("Password: " + acc.Password)
	color.Cyan("URL:" + acc.UrlString)
}
