package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Person struct {
	Id        int    `gorm:"column:id"`
	FirstName string `gorm:"column:firstname"`
	LastName  string `gorm:"column:lastname"`
}

func main() {
	dsn := "host=localhost user=mahdi password=123456 dbname=test port=5432 sslmode=disable client_encoding=UTF8"
	db, err3 := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err3 != nil {
		log.Fatalln(err3)
	}

	err2 := db.AutoMigrate(&Person{})
	if err2 != nil {
		log.Fatalln(err2)
	}

	//Insert
	//db.Table("person").Create(&Person{Id: 6, FirstName: "javad", LastName: "razavi"})

	//Read
	var person Person
	db.Table("person").First(&person, 4)
	fmt.Println(person.Id)
	fmt.Println(person.FirstName)
	fmt.Println(person.LastName)
}
