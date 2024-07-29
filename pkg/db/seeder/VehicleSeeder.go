package seeder

import (
	"log"

	"github.com/jaswdr/faker"
	"gorm.io/gorm"

	"github.com/argaputra12/car-repair/pkg/models"
)

func SeedVehicles(db *gorm.DB) {
	fake := faker.New()

	for i := 0; i < 10; i++ {
		vehicle := models.Vehicle{
			Make:   fake.Car().Maker(),
			Model:  fake.Car().Model(),
			Year:   fake.Time().Year(),
			Plate:  fake.Car().Plate(),
			UserID: uint(i + 1),
		}
		db.Create(&vehicle)
	}
	log.Println("Vehicles Seeded")
}
