package controllers

import (
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetActividadById(ctx *gin.Context) {
	actividadIDString := ctx.Param("id")

	actID, err := strconv.Atoi(actividadIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	actDep := services.GetActividadById(actID)

	ctx.JSON(http.StatusOK, actDep)
}

// GET /actividades/?q=spinning
func GetActividades(ctx *gin.Context) {
	q := ctx.Query("q")
	acts, err := services.BuscarActividades(q)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching actividades"})
		return
	}
	ctx.JSON(http.StatusOK, acts)
}

// POST /inscripcion   body: { "userId": 1, "actividadId": 3 }
type inscripcionBody struct {
	UserID      uint `json:"userId"`
	ActividadID uint `json:"actividadId"`
}

func Inscribirse(ctx *gin.Context) {
	var body inscripcionBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.InscribirUsuario(body.UserID, body.ActividadID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"msg": "Inscripcion exitosa"})
}

func MisActividades(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id inválido"})
		return
	}
	acts, err := services.ActividadesDeUsuario(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching actividades"})
		return
	}
	ctx.JSON(http.StatusOK, acts)
}
