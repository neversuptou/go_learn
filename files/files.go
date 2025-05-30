package files

import (
	"learnGO/utils"
	"os"
)

func SaveFile(content []byte, fileName string) {
	file, err := os.Create(fileName)
	utils.CheckError(err)

	defer file.Close()
	_, err = file.Write(content)
	utils.CheckError(err)

	return
}
func ReadFile(name string) ([]byte, error) {
	_, err := os.Open(name)
	utils.CheckError(err)

	data, err := os.ReadFile("data.json")
	utils.CheckError(err)

	return data, nil
}
