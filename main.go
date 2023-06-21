package main

import (
	"fmt"
	"gorm.io/driver/postgres" // or mysql or sqlite or sqlserver
	"gorm.io/gorm"
	"log"
)

// Engine is a type that represents an engine
type Engine struct {
	gorm.Model
	Power int `gorm:"column:power"` // the power of the engine in horsepower
}

// Car is a type that represents a car
type Car struct {
	gorm.Model
	EModel string  `gorm:"column:model"`               // the model of the car
	Engine *Engine `gorm:"column:engine_id default:1"` // the engine of the car
}

func main() {
	// open a connection to your database
	dsn := "host=localhost user=mahdi password=123456 dbname=test port=5432 sslmode=disable client_encoding=UTF8"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // change postgres to your driver name
	if err != nil {
		log.Fatal(err)
	}

	// create the tables for Car and Engine structs
	err = db.AutoMigrate(&Car{}, &Engine{})
	if err != nil {
		log.Fatal(err)
	}

	// create an engine
	e := &Engine{Power: 200}

	// create a car with the engine
	c := &Car{EModel: "Ford Mustang", Engine: e}

	// insert the car and its engine into the database
	res := db.Create(c)
	fmt.Println("Error: ", res.Error)
}
