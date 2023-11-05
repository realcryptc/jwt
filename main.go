package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	// other packages you may need
)

func GenerateToken() (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload) for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = "john.doe"
	claims["uuid"] = uuid.New()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate the token string
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func main() {
	token, err := GenerateToken()
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("Generated token:", token)
}
