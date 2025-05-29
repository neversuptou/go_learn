package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login    string
	password string
	url      string
}

func main() {
	fmt.Println(generatePassword(10))
	login := promptData("Enter your login: ")
	password := promptData("Enter your password: ")
	url := promptData("Enter URL: ")

	acc := account{
		login:    login,
		password: password,
		url:      url,
	}

	outputPassword(&acc)

}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc *account) {
	fmt.Println(acc.login, acc.password, acc.url)
}

func generatePassword(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789@-_")
	arr := make([]rune, n)

	for i := range arr {
		arr[i] = letters[rand.IntN(len(letters))]
	}

}
