package config

import (
	"fmt"
	"presensee_project/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "Kuripan69",
		"DB_Port":     "3306",
		"DB_Host":     "localhost",
		"DB_Name":     "projects",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
	return DB
}

func InitMigrate() {
	DB.AutoMigrate(&model.User{}, &model.Mahasiswa{})
}
