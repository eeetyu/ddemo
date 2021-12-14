package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MyDb struct {
	IsInit bool
	*gorm.DB
}

var Db MyDb

func InitDb() {
	dsn := "root:921110895@tcp(127.0.0.1:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Db.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Db.IsInit = true
}
