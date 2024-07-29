package seeder

import (
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
			Role:     "user",
			Phone:    fake.Phone().Number(),
			Address:  fake.Address().Address(),
		}
		db.Create(&user)
	}

	// Create Admin
	admin := models.User{
		Name:     "Admin",
		Email:    "admin@admin.com",
		Password: "password",
		Role:     "admin",
		Phone:    fake.Phone().Number(),
		Address:  fake.Address().Address(),
	}

	db.Create(&admin)

	log.Println("Users Seeded")
}
