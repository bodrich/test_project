package structures

// структура для записи в БД
type WritePerson struct {
	FirstName string // имя
	SurName   string // фамилия
	Sex       byte   // пол
}

// структура для чтения записи из БД и отдачи ее пользователю
type ReadPerson struct {
	Id        int64  // айди записи в БД
	FirstName string // имя
	SurName   string // фамилия
	Sex       byte   // пол
}

// структура для конфига
type Config struct {
	DatabaseHost     string // адрес базы данных
	DatabaseUser     string // имя пользователя
	DatabasePassword string // пароль
	DatabaseName     string // имя бд
	WebHost          string // адрес и порт, на котором будет работать веб-сервер
}
