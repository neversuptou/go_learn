package files

import (
	"learnGO/utils"
	"os"
)

type JsonDB struct {
	fileName string
}

func NewJsonDB(name string) *JsonDB {
	return &JsonDB{
		fileName: name,
	}
}

func (db *JsonDB) Write(content []byte) {
	file, err := os.Create(db.fileName)
	utils.CheckError(err)

	defer file.Close()
	_, err = file.Write(content)
	utils.CheckError(err)

	return
}
func (db *JsonDB) Read() ([]byte, error) {
	data, err := os.ReadFile(db.fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}
