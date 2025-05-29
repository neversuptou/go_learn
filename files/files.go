package files

import (
	"fmt"
	"learnGO/utils"
	"os"
)

func SaveFile(content []byte, fileName string) {
	file, err := os.Create(fileName)
	utils.CheckError(err)

	defer file.Close()
	_, err = file.Write(content)
	utils.CheckError(err)

	fmt.Println("File saved")

	return
}
func ReadFile() {
	file, err := os.Open("data.json")
	utils.CheckError(err)

	defer file.Close()
	data, err := os.ReadFile("data.json")
	utils.CheckError(err)

	fmt.Println(string(data))
}
