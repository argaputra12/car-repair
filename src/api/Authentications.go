package api

import (
	"net/http"

	"github.com/argaputra12/car-repair/pkg/db"
	"github.com/argaputra12/car-repair/pkg/models"
	"github.com/argaputra12/car-repair/pkg/utils"
)

// Login get page
func Login(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "login", nil)
}

// Register get page
func Register(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "register", nil)
}

// LoginPost post page
func LoginPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user := models.User{}

	db := db.Init()
	err := db.Where("email = ?", email).First(&user).Error

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "User not found")
		return
	}

	if user.Password != password {
		utils.RespondWithError(w, http.StatusBadRequest, "Password is incorrect")
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}

// RegisterPost post page
func RegisterPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user := models.User{
		Email:    email,
		Password: password,
		Role:     "user",
	}
	// check if email already exists
	existingUser := models.User{}

	db := db.Init()
	err := db.Where("email = ?", email).First(&existingUser).Error
	if err == nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Email already exists")
		return
	}
	
	err = db.Create(&user).Error

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to create user")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "User created successfully"})
}
