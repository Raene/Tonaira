package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

//Config type contains Router
type Config struct {
	Db  *gorm.DB
	Val *validator.Validate
}

//Init sets up the needed variables required for the server app to run
func Init(db *gorm.DB, val *validator.Validate) *Config {
	config := Config{}
	config.Db = db
	config.Val = val

	return &config
}
