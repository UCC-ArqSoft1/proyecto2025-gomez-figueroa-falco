package controllers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /inscripcion   body: { "userId": 1, "actividadId": 3 }
// Inscribirse maneja la inscripción de un usuario a un horario específico.
// El cuerpo de la solicitud debe contener el ID del usuario y el ID del horario.
type inscripcionBody struct {
	UserID      uint   `json:"userId" binding:"required"`
	Dia         string `json:"dia" binding:"required"`
	HorarioID   uint   `json:"horarioId" binding:"required"`
	ActividadID uint   `json:"actividadId" binding:"required"`
}

func Inscribirse(ctx *gin.Context) {
	var body inscripcionBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := services.InscribirUsuario(body.UserID, body.HorarioID); err != nil {
		switch err.Error() {
		case "usuario no encontrado":
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Usuario no existe"})
		case "horario no encontrado":
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Horario no existe"})
		case "usuario ya inscrito en este horario":
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ya estás inscripto"})
		case "cupo completo":
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cupo completo"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error inscribiendo"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"msg": "Inscripcion exitosa"})
}
