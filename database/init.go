package database

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"project/config"
	"project/logger"
)

// здесь храним подготовленный запрос для записи в БД
var PutRequest *sql.Stmt
// а здесь подготовленный запрос для чтения из БД
var SelectRequest *sql.Stmt


// Здесь инициализируем подключение к БД и подготовка запросов
// Вызывается один раз при запуске приложения
func InitDB() {
	logger.WriteLog("Start init database")
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable",
			config.Config.DatabaseHost, config.Config.DatabaseUser, config.Config.DatabaseName,
			config.Config.DatabasePassword))
	if err != nil {
		log.Fatalln(err)
	}

	PutRequest, err = db.Prepare("INSERT INTO person(first_name, sur_name, sex) VALUES($1, $2, $3) RETURNING id")
	if err != nil {
		log.Fatalln(err)
	}

	SelectRequest, err =  db.Prepare("SELECT id, first_name, sur_name, sex FROM person WHERE id=$1")
	if err != nil {
		log.Fatalln(err)
	}

	logger.WriteLog("Init database finished")
}
