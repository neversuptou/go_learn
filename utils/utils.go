package utils

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func PromptData[typeOfSlice any](prompts []typeOfSlice) string {
	reader := bufio.NewReader(os.Stdin)
	for i, prompt := range prompts {
		if i == len(prompts)-1 {
			fmt.Printf("%s: ", prompt)
		} else {
			fmt.Println(prompt)
		}
	}
	res, _ := reader.ReadString('\n')
	return strings.TrimSpace(res)
}

func CheckError(err error) {
	if err != nil {
		PrintError(err.Error())
	}
}

func PrintError(value any) {
	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки: %d", t)
	case error:
		color.Red(t.Error())
	default:
		color.Red("Unknown error")
	}

}
