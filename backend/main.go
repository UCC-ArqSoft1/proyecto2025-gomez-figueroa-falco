package main

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/actividades", controllers.GetActividades)
	router.GET("/actividad/:id", controllers.GetActividadById)
	router.POST("/inscripcion", controllers.Inscribirse)
	router.GET("/misActividades/:userId", controllers.MisActividades)

	router.POST("/login", controllers.Login)

	router.Run(":8080")

}
