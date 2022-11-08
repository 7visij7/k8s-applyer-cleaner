package config

import (
	"os"
    "io/ioutil"
	"gopkg.in/yaml.v3"
    "smib-applyers-cleaner/pkg/errors"
)

func GetCreds() map[string]string{
	jsonFile, err := os.Open("config.yaml")
	errors.CheckError(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data map[string]string
	err = yaml.Unmarshal(byteValue, &data)
	errors.CheckError(err)

    return data
}