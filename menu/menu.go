package menu

import (
	"learnGO/account"
	"learnGO/files"
	"learnGO/utils"
)

var menu = map[string]func(*account.VaultWithDB){
	"1": account.CreateAccount,
	"2": account.FindAccountWithUrl,
	"3": account.FindAccountWithLogin,
	"4": account.DeleteAccount,
}

func ChooseMenu() (breaked bool) {
	//vault := account.NewVault(cloud.NewCloudDB("https://vk.com"))
	vault := account.NewVault(files.NewJsonDB("data.json"))
	for {
		variant := []string{"1. Create account",
			"2. Find account by url",
			"3. Find account by login",
			"4. Delete",
			"5. Exit",
			"Choose option ",
		}
		choosed := utils.PromptData(variant)
		menuFunc := menu[choosed]
		if menuFunc == nil {
			return true
		}
		menuFunc(vault)
	}
}
