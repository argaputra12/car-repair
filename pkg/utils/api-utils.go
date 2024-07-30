package utils

import (
	"encoding/json"
	"net/http"
	"text/template"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// RenderTemplate render template
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpl = "templates/" + tmpl + ".html"
	t, _ := template.ParseFiles(tmpl)
	t.Execute(w, data)
}

// RespondWithError respond with error
func RespondWithError(w http.ResponseWriter, code int, message string) {

	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON respond with json
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// GenerateToken generate token
func GenerateToken(userID uint, role string) (string, error) {
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"role":   role,
		"time":   time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with the secret
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
