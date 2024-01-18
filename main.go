package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(
		"postgres",
		"host=localhost port=5432 user=postgres dbname=test password=postgres sslmode=disable",
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected")
	db.AutoMigrate(&Contact{}, &Task{}, &Reminder{})
}
