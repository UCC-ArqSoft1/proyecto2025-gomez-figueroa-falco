package services

import (
	"backend/clients"
	"fmt"
)

func Login(username string, password string) {
	// Get user by
	user := clients.GetUserByUsername(username)
	fmt.Println("Usuario encontrado:", user)
	if user.Id == 0 {
		fmt.Println("Usuario no encontrado")
		return
	}
}
