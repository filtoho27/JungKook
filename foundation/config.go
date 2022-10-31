package foundation

import (
	"io/ioutil"
	"log"
	"os"
)

func GetConfig(path string) (result []byte, err error) {
	configPath := os.Getenv("CONFIGPATH")
	jsonFile, err := os.Open(configPath + path)
	if err != nil {
		log.Printf("%+v", err)
	}
	defer jsonFile.Close()
	result, _ = ioutil.ReadAll(jsonFile)
	return
}
