package menu

import (
	"learnGO/account"
	"learnGO/encrypter"
	"learnGO/files"
	"learnGO/utils"
)

var menu = map[string]func(*account.VaultWithDB){
	"1": account.CreateAccount,
	"2": account.FindAccountWithUrl,
	"3": account.FindAccountWithLogin,
	"4": account.DeleteAccount,
}

var menuVariants = []string{
	"1. Create account",
	"2. Find account by url",
	"3. Find account by login",
	"4. Delete",
	"5. Exit",
}

func ChooseMenu() (breaked bool) {
	//vault := account.NewVault(cloud.NewCloudDB("https://vk.com"))
	vault := account.NewVault(files.NewJsonDB("data.vault"), *encrypter.NewEncrypter())
	for {
		variant := utils.PromptData(menuVariants...)
		menuFunc := menu[variant]
		if menuFunc == nil {
			return true
		}
		menuFunc(vault)
	}
}
