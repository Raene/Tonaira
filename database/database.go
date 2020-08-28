package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	//import driver for effect
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

//Init sets up the database
func Init() *gorm.DB {
	viper.SetConfigFile(".env")
	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// To get the value from the config file using key
	// viper package read .env
	viperUser := viper.Get("POSTGRES_USER")
	viperPassword := viper.Get("POSTGRES_PASSWORD")
	viperDb := viper.Get("POSTGRES_DB")
	viperHost := viper.Get("POSTGRES_HOST")
	viperPort := viper.Get("POSTGRES_PORT")

	// https://gobyexample.com/string-formatting
	prosgretConname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viperHost, viperPort, viperUser, viperDb, viperPassword)

	db, err := gorm.Open("postgres", prosgretConname)
	if err != nil {
		panic("Failed to connect to database!")
	}
	return db
}
