package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pajaro.com/curso-go/sistema"
)

var Database = func() (db *gorm.DB)  {
	mysqluser := sistema.GoDotEnvVariable("MYSQL_USER")
	if mysqluser == "" {
		mysqluser = "curso-go"
	}

	mysqlpassword := sistema.GoDotEnvVariable("MYSQL_PASSWORD")
	if mysqlpassword == "" {
		mysqlpassword = "curso-go"
	}

	mysqldb := sistema.GoDotEnvVariable("MYSQL_DATABASE")
	if mysqldb == "" {
		mysqldb = "curso-go"
	}

	var url string = mysqluser + ":" + mysqlpassword + "@tcp(db:3306)/" + mysqldb + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{}); if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Conectado a la base de datos " + mysqldb)
		return db
	}
}()