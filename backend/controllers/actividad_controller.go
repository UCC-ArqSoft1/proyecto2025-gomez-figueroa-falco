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
