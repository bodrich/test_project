package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"project/config"
	"project/logger"
)

func SetupServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.PUT("/insert", Put)
	r.GET("/get/:id", Get)

	return r
}

// инициализация веб-сервера
func InitServer() {
	logger.WriteLog("Start init web server")
	err := SetupServer().Run(config.Config.WebHost)
	if err != nil {
		log.Fatal(err)
	}
}
