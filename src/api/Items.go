package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/argaputra12/car-repair/pkg/db"
	"github.com/argaputra12/car-repair/pkg/models"
	"github.com/argaputra12/car-repair/pkg/utils"
)

// Get all items
func GetItems(w http.ResponseWriter, r *http.Request) {
	db := db.Init()
	items := []models.Item{}
	db.Find(&items)

	utils.RespondWithJSON(w, http.StatusOK, items)
}

// Get item by id
func GetItem(w http.ResponseWriter, r *http.Request) {
	db := db.Init()
	item := models.Item{}
	id := r.URL.Query().Get("id")
	db.First(&item, id)

	utils.RespondWithJSON(w, http.StatusOK, item)
}

// Create Item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.PostFormValue("name")
	price := r.PostFormValue("price")
	quantity := r.PostFormValue("quantity")

	// Convert price to float64
	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		// Handle the error, e.g., log it, return an error response, etc.
		log.Println("Error parsing price:", err)
		return
	}

	// Convert quantity to int
	quantityInt, err := strconv.Atoi(quantity)
	if err != nil {
		// Handle the error, e.g., log it, return an error response, etc.
		log.Println("Error parsing quantity:", err)
		return
	}

	item := models.Item{
		Name:     name,
		Price:    priceFloat,
		Quantity: quantityInt,
	}

	db := db.Init()
	db.Create(&item)

	utils.RespondWithJSON(w, http.StatusOK, item)
}

// Update item
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id := r.PostFormValue("id")
	name := r.PostFormValue("name")
	price := r.PostFormValue("price")
	quantity := r.PostFormValue("quantity")

	// Convert price to float64
	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		// Handle the error, e.g., log it, return an error response, etc.
		log.Println("Error parsing price:", err)
		return
	}

	// Convert quantity to int
	quantityInt, err := strconv.Atoi(quantity)
	if err != nil {
		// Handle the error, e.g., log it, return an error response, etc.
		log.Println("Error parsing quantity:", err)
		return
	}

	db := db.Init()
	db.Model(&models.Item{}).Where("id = ?", id).Updates(models.Item{Name: name, Price: priceFloat, Quantity: quantityInt})

	utils.RespondWithJSON(w, http.StatusOK, nil)
	
}

// Delete item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id := r.PostFormValue("id")

	db := db.Init()
	db.Where("id = ?", id).Delete(&models.Item{})

	utils.RespondWithJSON(w, http.StatusOK, nil)
}