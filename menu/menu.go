package menu

import (
	"bufio"
	"fmt"
	"learnGO/account"
	"os"
	"strconv"
	"strings"
)

func ChooseMenu() (breaked bool) {
	reader := bufio.NewReader(os.Stdin)

	for {
		viewMenu()
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		key, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Пожалуйста, введите число")
			continue
		}

		switch key {
		case 1:
			account.CreateAccount()
		case 2:
			account.FindAccount()
		case 3:
			account.DeleteAccount()
		case 4:
			fmt.Println("До свидания!") //
			return true
		default:
			fmt.Println("Неверный выбор. Пожалуйста, выберите от 1 до 4")
		}
		return false
	}
}

func viewMenu() {
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	return
}
