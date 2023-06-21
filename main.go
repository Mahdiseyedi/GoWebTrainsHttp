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
	//db.Table("person").Create(&Person{Id: 5, FirstName: "kave", LastName: "Ahani"})

	//bulk-insert
	persons := []*Person{
		&Person{Id: 7, FirstName: "kevin", LastName: "anderson"},
		&Person{Id: 8, FirstName: "john", LastName: "wick"},
		&Person{Id: 9, FirstName: "alex", LastName: "jude"},
		&Person{Id: 10, FirstName: "rose", LastName: "peterson"},
	}
	//
	//persons := []Person{
	//	Person{Id: 11, FirstName: "kevin", LastName: "anderson"},
	//	Person{Id: 12, FirstName: "john", LastName: "wick"},
	//	Person{Id: 13, FirstName: "alex", LastName: "jude"},
	//	Person{Id: 14, FirstName: "rose", LastName: "peterson"},
	//}
	res := db.Table("person").Create(persons)
	fmt.Println(res)
}
