package main

import (
	"backend/controllers"

	"github.com/gin-contrib/cors"

	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/actividades", controllers.GetActividades)
	router.GET("/actividad/:id", controllers.GetActividadById)
	router.POST("/inscripcion", controllers.Inscribirse)
	router.GET("/misActividades/:userId", controllers.MisActividades)

	router.POST("/login", controllers.Login)

	router.Run(":8080")

}
