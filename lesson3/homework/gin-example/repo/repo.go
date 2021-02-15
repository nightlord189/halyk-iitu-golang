package repo

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//readFile Читает файл. расположенный в пути path и заполняет интерфейс v содержимым этого файла
func readFile(path string, v interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func writeFile(path string, v interface{}) error {
	bytes, err := json.MarshalIndent(v, "", "	")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, bytes, 0644)
	return err
}

// Init Инициализирует репозиторий
func Init() error {
	err := uploadUserStates()
	if err != nil {
		return err
	}
	err = uploadUsers()
	if err != nil {
		return err
	}
	err = initTokensStore()
	if err != nil {
		return err
	}
	return nil
}
