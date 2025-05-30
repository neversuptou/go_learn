package menu

import (
	"bufio"
	"fmt"
	"learnGO/account"
	"learnGO/files"
	"learnGO/utils"
	"os"
	"strconv"
	"strings"
)

func ChooseMenu() (breaked bool) {
	//vault := account.NewVault(cloud.NewCloudDB("https://vk.com"))
	vault := account.NewVault(files.NewJsonDB("data.json"))
	reader := bufio.NewReader(os.Stdin)
	for {
		viewMenu()
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		key, err := strconv.Atoi(input)
		if err != nil {
			utils.PrintError("Please, enter a number")
			continue
		}

		switch key {
		case 1:
			account.CreateAccount(vault)
		case 2:
			account.FindAccount(vault)
		case 3:
			account.DeleteAccount(vault)
		case 4:
			fmt.Println("Good Buy!") //
			return true
		default:
			fmt.Println("Wrong choice. Please, choose again")
		}
		return false
	}
}

func viewMenu() {
	fmt.Println("1. Create account")
	fmt.Println("2. Find account")
	fmt.Println("3. Delete account")
	fmt.Println("4. Exit")
	return
}
