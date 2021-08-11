package utils

import (
	"io/ioutil"
	"log"
	"os"
)

// it used to read file
func ReadFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("ReadFile Error: ", err)
		return nil, err
	}
	return data, nil
}

// it used to write data into file
func WriteFile(filename string, data []byte) bool {

	if _, err := os.Stat(filename); err != nil {
		_, fileErr := os.Create(filename)
		if fileErr != nil {
			log.Println("WriteFile Error: ", fileErr)
			return false
		}
	}
	err := ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		log.Println("WriteFile Error: ", err)
		return false
	}
	return true
}
