package services

import (
	"backend/dao"
	"backend/clients"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserId uint   `json:"userId"`
	Rol    string `json:"rol"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("jwtSecret")

func GenerateToken(UserId uint, rol string) (string, error) {
	claims := CustomClaims{
		UserId: UserId, // id del usuario
		Rol:    rol,    // rol del usuario
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24 horas de expiracion
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // fecha de emision
			Issuer:    "backend",                                          // emisor
			Subject:   "auth",                                             // asunto
			ID:        fmt.Sprintf("%d", UserId),                          // id del usuario

		},
	}
	//crear token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//firmar el token
	signed, err := token.SignedString([]byte(jwtKey)) // usar la clave secreta del entorno

	//valida el error
	if err != nil {
		return "", fmt.Errorf("error firmando JWT: %w", err)
	}
	//vuelve el token
	return signed, nil
}

func Login(username string, password string) (string, error) {
	// Get user by
	user, err := dao.GetUserByUsername(clients.DB, username)
	if err != nil {
		fmt.Println("Usuario no encontrado")
		return "", fmt.Errorf("usuario no encontrado")
	}
	fmt.Println("Usuario encontrado:", user)

	// Hash the password
	hashedPassword := sha256.Sum256([]byte(password))
	fmt.Print("Contraseña hasheada:", hashedPassword)
	if user.PasswordHash != fmt.Sprintf("%x", hashedPassword) {
		fmt.Println("Contraseña incorrecta")
		return "", fmt.Errorf("contraseña incorrecta")
	}
	// Generate JWT token
	token, err := GenerateToken(user.Id, user.Rol)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return "", err
	}
	// Return the token
	return token, nil
}
