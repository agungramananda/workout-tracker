package configs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func LoadDBConfig() *DBConfig {
	return &DBConfig{
		Host:     GetDotEnvVariable("DB_HOST"),
		Port:     GetDotEnvVariable("DB_PORT"),
		User:			GetDotEnvVariable("DB_USER"),
		Password: GetDotEnvVariable("DB_PASSWORD"),
		DBName	: GetDotEnvVariable("DB_NAME"),
	}
}

func ConnectDB() (*gorm.DB,error){
	dbconfig := LoadDBConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbconfig.Host, dbconfig.User, dbconfig.Password, dbconfig.DBName, dbconfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
