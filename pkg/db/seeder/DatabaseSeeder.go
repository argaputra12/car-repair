package seeder

import (
	"log"

	"github.com/argaputra12/car-repair/pkg/db"
)

func SeedDatabase() {
	db := db.Init()

	SeedUsers(db)
	SeedVehicles(db)
	SeedItem(db)

	log.Println("Migration has been processed")
}
