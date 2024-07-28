package seeder

import (
	"fmt"
	"log"

	"github.com/argaputra12/car-repair/pkg/models"

	"github.com/jaswdr/faker"
	"gorm.io/gorm"
)

func SeedItem(db *gorm.DB) {
	fake := faker.New()

	for i := 0; i < 10; i++ {
		item := models.Item{
			Name:     fake.Person().FirstName(),
			Price:    fake.Float64(7, 1000, 9999999),
			Quantity: fake.RandomNumber(2),
		}
		db.Create(&item)
	}
	fmt.Println("Items Seeded")
	log.Println("Items Seeded")
}
