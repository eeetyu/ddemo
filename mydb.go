package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MyDb struct {
	IsInit bool
	*gorm.DB
}

var DbClient MyDb

func (mydb *MyDb) InitDb() {
	dsn := "sun:aA921110895@tcp(111.62.122.237:3306)/sun_demo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	mydb.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	mydb.IsInit = true
}
