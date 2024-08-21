package api

import (
	"net/http"

	"github.com/argaputra12/car-repair/pkg/db"
	"github.com/argaputra12/car-repair/pkg/models"
	"github.com/argaputra12/car-repair/pkg/utils"

)

// Get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := db.Init()
	users := []models.User{}
	db.Find(&users)

	utils.RespondWithJSON(w, http.StatusOK, users)
}

// Get user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	db := db.Init()
	user := models.User{}
	id := r.URL.Query().Get("id")
	db.First(&user, id)

	utils.RespondWithJSON(w, http.StatusOK, user)
}

// Create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user := models.User{
		Email:    email,
		Password: password,
	}

	// check if email already exists
	db := db.Init()
	err := db.Where("email = ?", email).First(&user).Error
	if err == nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Email already exists")
		return
	}

	db.Create(&user)

	utils.RespondWithJSON(w, http.StatusOK, user)
}

// Update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id := r.PostFormValue("id")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user := models.User{
		Email:    email,
		Password: password,
	}

	db := db.Init()
	db.Model(&user).Where("id = ?", id).Updates(&user)

	utils.RespondWithJSON(w, http.StatusOK, user)
}

// Delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id := r.PostFormValue("id")

	user := models.User{}

	db := db.Init()
	db.Where("id = ?", id).Delete(&user)

	utils.RespondWithJSON(w, http.StatusOK, user)
}