package config

import (
	"db"
	"encoding/json"
	"fmt"
	"logger"
	"os"
)

type ApplicationConfig struct {
	DbConfig *db.DbConfig
	LogFile  *string
	Port     *int64
}

var AppConfig = ApplicationConfig{}

func Init(fileName string) bool {
	ret := true
	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&AppConfig)
	if err != nil {
		logger.Error("%s", err)
		fmt.Println("%s", err)
		ret = false
	}
	return ret
}

/*
func (appConfig *ApplicationConfig) SetConfigProperties() *db.DbConfig {
	file, err := ioutil.ReadFile("src/config/config.json")
	if err != nil {
		logger.Debug("%s", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &appConfig)
	return appConfig.DbConfig
}
*/
