package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	secretKey := []byte("mySecretKey")
	claims := jwt.MapClaims{
		"userID": 123,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return
	}
	fmt.Println("Generated Token:", signedToken)

	validateToken(signedToken)
}

func validateToken(signedToken string) {
	secretKey := []byte("mySecretKey")

	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err)
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Token is valid: Claims")
		fmt.Println("UserID:", claims["userID"])
		fmt.Println("Expires at:", claims["exp"])
	} else {
		fmt.Println("Invalid token")
	}
}
