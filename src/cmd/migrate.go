package main

import (
	"fmt"
	"log"

	"github.com/argaputra12/car-repair/pkg/db"
	"github.com/argaputra12/car-repair/pkg/models"
	"github.com/argaputra12/car-repair/pkg/db/seeder"
)

func main() {
	db := db.Init()

	// Delete Table if exists
	db.Migrator().DropTable(&models.User{}, &models.Vehicle{}, &models.Item{}, &models.Service{}, &models.ServiceItem{}, &models.ServiceRecord{}, &models.Payment{})

	// Migrate Table
	err := db.AutoMigrate(&models.User{}, &models.Vehicle{}, &models.Item{}, &models.Service{}, &models.ServiceItem{}, &models.ServiceRecord{}, &models.Payment{})

	if err != nil {
		log.Fatal(err)
	}

	seeder.SeedDatabase()

	fmt.Println("Migration has been processed")
	log.Println("Migration has been processed")
}
