package main

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/actividad/:id", controllers.GetActividadById)
	router.POST("/login", controllers.Login)
	router.Run(":8080")

}
