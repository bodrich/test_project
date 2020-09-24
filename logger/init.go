package logger

import (
	"log"
	"os"
	"time"
)

func InitLogger() {
	file, err := os.OpenFile("logs/" + time.Now().Format("01-02-2006") + ".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	WriteLog("Logger init finished")
}

// запись в лог
func WriteLog(format string, args ...interface{}) {
	log.Printf(format + "\n", args...)
}

