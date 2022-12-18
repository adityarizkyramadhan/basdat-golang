package main

import (
	"fmt"
	"os"

	"basdat-golang/model"
	querysql "basdat-golang/repository/query_sql"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
	driverDb, err := ReadEnvDatabase()
	if err != nil {
		panic(err.Error())
	}
	db, err := MakeConnection(driverDb)
	if err != nil {
		panic(err.Error())
	}
	data := &model.Pelanggan{
		Name:    "Aditya Rizky Ramadhan",
		Address: "Sidoarjo",
	}
	err = db.Exec(querysql.AddDataTable(data)).Error
	if err != nil {
		panic(err.Error())
	}
}

func ReadEnvDatabase() (DriverSupabase, error) {
	return DriverSupabase{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
	}, nil
}

type DriverSupabase struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
}

func MakeConnection(data DriverSupabase) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", data.User, data.Password, data.Host, data.Port, data.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
