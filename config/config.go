package config

import "github.com/jinzhu/gorm"

//Config type contains Router
type Config struct {
	Db *gorm.DB
}

//Init sets up the needed variables required for the server app to run
func Init(db *gorm.DB) *Config {
	config := Config{}
	config.Db = db

	return &config
}
