package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"project/logger"
	"project/structures"
)

var Config structures.Config

func InitConfig() {
	logger.WriteLog("Start init config")
	file, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = json.Unmarshal(file, &Config); err != nil {
		log.Fatal(err.Error())
	}

	logger.WriteLog("Init config finished")
}