package main

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/actividad/:id", controllers.GetActividadById)

	router.Run(":8080")
}
