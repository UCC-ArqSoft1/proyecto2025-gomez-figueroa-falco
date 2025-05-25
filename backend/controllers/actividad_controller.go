package controllers

import (
	"backend/dto"
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

// GET /misActividades/:userId
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

func CrearActividad(ctx *gin.Context) {
	//Validar rol desde el token
	rol, ok := ctx.Get("rol")
	if !ok || rol != "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "No autorizado"})
		return
	}

	var input dto.ActividadConHorarioRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := services.CrearActividadConHorario(input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"msg": "Actividad creada"})
}

func EditarActividad(ctx *gin.Context) {
	rol, ok := ctx.Get("rol")
	if !ok || rol != "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "No autorizado"})
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id inválido"})
		return
	}

	var input dto.ActividadConHorarioRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := services.ActualizarActividad(uint(id), input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Actividad editada"})
}

func EliminarActividad(ctx *gin.Context) {
	rol, ok := ctx.Get("rol")
	if !ok || rol != "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "No autorizado"})
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id inválido"})
		return
	}
	if err := services.EliminarActividad(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Actividad eliminada"})
}
