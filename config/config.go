package config

import (
	"fmt"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var appOnce = sync.Once{}
var dbOnce = sync.Once{}

// Application holds application config
type Application struct {
	Env             string `mapstructure:"ENV"`
	ServerAddress   string `mapstructure:"SERVER_ADDRESS"`
	IsLoggingToFile bool   `mapstructure:"IS_LOGGING_TO_FILE"`
	LogFilePath     string `mapstructure:"LOG_FILE_PATH"`
}

// DB holds database config
type DB struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	DBName   string `mapstructure:"DB_NAME"`
}

var appConfig *Application
var dbConfig *DB

func loadApp() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env not found, that's okay!")
	}

	viper.AutomaticEnv()

	appConfig = &Application{
		Env:             viper.GetString("ENV"),
		ServerAddress:   viper.GetString("SERVER_ADDRESS"),
		IsLoggingToFile: viper.GetBool("IS_LOGGING_TO_FILE"),
		LogFilePath:     viper.GetString("LOG_FILE_PATH"),
	}

	return nil
}

func loadDB() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env not found, that's okay!")
	}

	viper.AutomaticEnv()

	dbConfig = &DB{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		DBName:   viper.GetString("DB_NAME"),
	}

	return nil
}

// GetApp returns application config
func GetApp() *Application {
	appOnce.Do(func() {
		loadApp()
	})

	return appConfig
}

// GetDB returns database config
func GetDB() *DB {
	dbOnce.Do(func() {
		loadDB()
	})

	return dbConfig
}
