package main

import (
	"project/config"
	"project/database"
	"project/logger"
	"project/web"
)

func main() {
	logger.WriteLog("Start application...")
	config.InitConfig()
	logger.InitLogger()
	database.InitDB()
	web.InitServer()
}
