package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}

func main() {
	dsn := "host=localhost user=mahdi password=123456 dbname=test port=5432 sslmode=disable client_encoding=UTF8"
	db, err3 := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err3 != nil {
		log.Fatalln(err3)
	}

	err2 := db.AutoMigrate(&User{}, &CreditCard{})
	if err2 != nil {
		log.Fatalln(err2)
	}

	//Associations Insert
	res := db.Create(&User{
		Name:       "mahdi",
		CreditCard: CreditCard{Number: "6037998023134025"},
	})
	fmt.Println(res.Error)
}
