package configs

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Config struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func Load() *Config {
	cfg := &Config{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}
	if cfg.DB_HOST == "" {
		log.Fatal("缺少环境变量 DB_HOST")
	}
	if cfg.DB_PORT == "" {
		log.Fatal("缺少环境变量 DB_PORT")
	}
	if cfg.DB_USER == "" {
		log.Println("缺少环境变量 DB_USER")
	}
	if cfg.DB_PASSWORD == "" {
		log.Println("缺少环境变量 DB_PASSWORD")
	}
	if cfg.DB_NAME == "" {
		log.Println("缺少环境变量 DB_NAME")
	}
	return cfg
}

func InitDB() error {
	config := Load()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
