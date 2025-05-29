package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"learnGO/files"
	"learnGO/utils"
	"math/rand/v2"
	"net/url"
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

func (acc AccountStruct) OutputData() {
	fmt.Printf("Аккаунт успешно сохранён: \nЛогин: %s\nПароль: %s\nURL: %s\n", acc.Login, acc.Password, acc.UrlString)
}

func (acc *AccountStruct) generatePassword(n int) {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789@-_")
	arr := make([]rune, n)

	for i := range arr {
		arr[i] = letters[rand.IntN(len(letters))]
	}
	acc.Password = string(arr)
}

func (acc *AccountStruct) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

//func (acc *AccountStruct) FromBytes() ([]byte, error) {
//	file, err := json.Unmarshal(acc)
//	if err != nil {
//		return nil, err
//	}
//	return file, nil
//}

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

func FindAccount() {
	//зАпрашиваем URL
	//делаем метод, который по url находит аккаунты
	//strings contained
	// вывод
	urlString := utils.PromptData("Напишите url доступы которого хотите найти: ")
	findAccountByURL(urlString)
}

func DeleteAccount() *AccountStruct {
	urlString := utils.PromptData("Напишите url доступы которого хотите удалить: ")

	found := false
	for i, v := range accountArray {
		if v.UrlString == urlString {
			accountArray = append(accountArray[:i], accountArray[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Аккаунт не найден")
	} else {
		fmt.Println("Аккаунт удален")
	}

	return nil
}

func CreateAccount() {
	login := utils.PromptData("Enter your login: ")
	password := utils.PromptData("Enter your Password: ")
	urlString := utils.PromptData("Enter your URL: ")

	// This creates the account and handles password generation if needed
	dataAccount, err := NewAccount(login, password, urlString)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := dataAccount.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}
	files.SaveFile(file, "data.json")

	dataAccount.OutputData()
}
