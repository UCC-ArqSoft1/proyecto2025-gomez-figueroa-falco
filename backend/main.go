package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Bienvenido al backend de tu proyecto!")
}

func main() {
	http.HandleFunc("/", homeHandler)

	port := "8080"
	fmt.Printf("Servidor iniciado en http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}