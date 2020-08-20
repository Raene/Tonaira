package config

import "database/sql"

//Config type contains Router
type Config struct {
	Db *sql.DB
}

//Init sets up the needed variables required for the server app to run
func Init(db *sql.DB) *Config {
	config := Config{}
	config.Db = db

	return &config
}
