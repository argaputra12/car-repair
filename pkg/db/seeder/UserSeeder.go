package seeder

import (
	"fmt"
	"log"

	"github.com/jaswdr/faker"
	"gorm.io/gorm"

	"github.com/argaputra12/car-repair/pkg/models"
)

func SeedUsers(db *gorm.DB) {
	fake := faker.New()

	for i := 0; i < 10; i++ {
		user := models.User{
			Name:     fake.Person().Name(),
			Email:    fake.Internet().Email(),
			Password: "password",
			Phone:    fake.Phone().Number(),
			Address:  fake.Address().Address(),
		}
		db.Create(&user)
	}
	fmt.Println("Users Seeded")
	log.Println("Users Seeded")
}
