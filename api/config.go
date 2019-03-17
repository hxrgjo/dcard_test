package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB struct {
		Host     string
		User     string
		Password string
		DBName   string
	}
	Port      int
	JWTSecret string
}

// Initial config from config file
func readConfig() (config Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	return Config{
		DB: struct {
			Host     string
			User     string
			Password string
			DBName   string
		}{
			Host:     viper.GetString("db.host"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
			DBName:   viper.GetString("db.dbname"),
		},
		Port:      viper.GetInt("port"),
		JWTSecret: viper.GetString("jwtSecret"),
	}, nil
}
