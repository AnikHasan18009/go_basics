package configuration

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	PORT        string `json:"port"`
	DB_PORT     string `json:"dbport"`
	HOST        string `json:"host"`
	DB_NAME     string `json:"dbname"`
	DB_USER     string `json:"user"`
	DB_PASSWORD string `json:"password"`
	SSL         string `json:"sslmode"`
	SECRET_KEY  string `json:"secretkey"`
}

var config Config

func LoadConfig(fileName string) (Config, error) {
	var currentConfig Config
	configFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error reading jsonfile:", err)
		return currentConfig, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&currentConfig)
	configFile.Close()

	if err != nil {
		log.Fatal("Error parsing json data:", err)
		return currentConfig, err
	}
	config = currentConfig
	return currentConfig, err
}

func GetConfig() Config { return config }
