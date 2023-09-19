package global

type Config struct {
	App App
	DB  DB
	JWT JWT
}

type App struct {
	Port int
	Env  string
}

type DB struct {
	Host     string
	User     string
	Password string
	Database string
}

type JWT struct {
	Secret string
}
