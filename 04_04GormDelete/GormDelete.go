package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

//https://gorm.io/docs/models.html#Fields-Tags

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

	//Read
	//var person Person
	//db.Table("person").First(&person, "firstname = 'arman'")
	////Query :
	//// SELECT * FROM "person" WHERE firstname = 'arman' ORDER BY "person"."id" LIMIT 1
	//fmt.Println(person.Id)
	//fmt.Println(person.FirstName)
	//fmt.Println(person.LastName)

	//Update
	//db.Table("person").Model(&person).Update("lastname", "akbari")
	//db.Table("person").Model(&person).Updates(Person{FirstName: "mohammad", LastName: "seyedi"})
	//fmt.Println(person.Id)
	//fmt.Println(person.FirstName)
	//fmt.Println(person.LastName)

	//Delete, both worked !
	var person Person
	//db.Table("person").Delete(&person, 5)
	db.Table("person").Model(&person).Delete(&person, 5)
	fmt.Println(person.Id)
	fmt.Println(person.FirstName)
	fmt.Println(person.LastName)
}
