package account

import "fmt"

func findAccountByURL(urlString string) {
	found := false
	for i, v := range accountArray {
		if v.UrlString == urlString {
			fmt.Printf("Найден аккаунт:\nID:%d \nЛогин: %s\nПароль: %s\nURL: %s\n",
				i+1, v.Login, v.Password, v.UrlString)
			found = true
		}
	}
	if !found {
		fmt.Println("Аккаунт не найден")
	}
}
