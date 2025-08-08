package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/geekAshish/DriveDesk/models"
	"github.com/geekAshish/DriveDesk/middleware"
	"github.com/golang-jwt/jwt/v5"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	valid := (credentials.UserName == "admin" && credentials.Password == "admin123")

	if !valid {
		http.Error(w, "Incorrect user name password", http.StatusUnauthorized)
		return
	}

	tokenString, err := GenerateToken(credentials.UserName)
	if err != nil {
		http.Error(w, "Unable to generate token", http.StatusInternalServerError)
		fmt.Println("Unable to generate token", err)
	}

	response := map[string]string{"token": tokenString}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GenerateToken(userName string) (string, error) {
	expiration := time.Now().Add(24 * time.Hour)

	claims := &middleware.Claims{
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("some_value"))
	if err != nil {
		return "", err
	}

	return signedToken, nil

}
