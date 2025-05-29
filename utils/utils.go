package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptData(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	res, _ := reader.ReadString('\n')
	return strings.TrimSpace(res)
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
