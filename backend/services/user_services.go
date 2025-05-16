package services

import (
	"backend/clients"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId uint) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24 horas de expiracion
		IssuedAt:  jwt.NewNumericDate(time.Now()),                     // fecha de emision
		Issuer:    "backend",                                          // emisor
		Subject:   "auth",                                             // asunto
		ID:        fmt.Sprintf("%d", userId),                          // id del usuario

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//firmar el token
	tokenString, err := token.SignedString([]byte("secret"))

	//valida el error
	if err != nil {
		return "", err
	}
	//vuelve el token
	return tokenString, nil
}

func Login(username string, password string) (string, error) {
	// Get user by
	user := clients.GetUserByUsername(username)
	fmt.Println("Usuario encontrado:", user)
	if user.Id == 0 {
		fmt.Println("Usuario no encontrado")
		return "", fmt.Errorf("usuario no encontrado")
	}

	// Hash the password
	hashedPassword := sha256.Sum256([]byte(password))
	fmt.Print("Contraseña hasheada:", hashedPassword)
	if user.PasswordHash != fmt.Sprintf("%x", hashedPassword) {
		fmt.Println("Contraseña incorrecta")
		return "", fmt.Errorf("contraseña incorrecta")
	}
	// Generate JWT token
	token, err := GenerateToken(user.Id)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return "", err
	}
	// Return the token
	return token, nil

}
