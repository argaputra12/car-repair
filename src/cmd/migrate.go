package main

import (
	"log"

	"github.com/argaputra12/car-repair/pkg/db"
	"github.com/argaputra12/car-repair/pkg/db/seeder"
	"github.com/argaputra12/car-repair/pkg/models"
)

func main() {
	db := db.Init()

	// Create Database if not exists
	db.Exec("CREATE DATABASE IF NOT EXISTS car_repair")

	// Delete Table if exists
	db.Migrator().DropTable(&models.User{}, &models.Vehicle{}, &models.Item{}, &models.Service{}, &models.ServiceItem{}, &models.ServiceRecord{}, &models.Payment{})

	// Migrate Table
	err := db.AutoMigrate(&models.User{}, &models.Vehicle{}, &models.Item{}, &models.Service{}, &models.ServiceItem{}, &models.ServiceRecord{}, &models.Payment{})

	if err != nil {
		log.Fatal(err)
	}

	seeder.SeedDatabase()

	log.Println("Migration has been processed")
}
