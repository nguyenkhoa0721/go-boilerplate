package config

type Postgres struct {
	Driver   string
	Host     string
	Port     uint
	Database string
	Username string
	Password string
}

type Redis struct {
	Host     string
	Port     uint
	Password string
}
