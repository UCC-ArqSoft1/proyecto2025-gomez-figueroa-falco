package controllers

import (
	"backend/dto"
	"backend/services"

	"encoding/json"
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
	rol, ok := ctx.Get("rol")
	if !ok || rol != "ADMIN" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "No autorizado"})
		return
	}

	nombre := ctx.PostForm("nombre")
	descripcion := ctx.PostForm("descripcion")
	categoria := ctx.PostForm("categoria")
	profesor := ctx.PostForm("profesor")
	cupoTotalS := ctx.PostForm("cupo_total")

	// Leer horarios como JSON
	horariosStr := ctx.PostForm("horarios")
	var horariosInput []dto.HorarioRequest
	if err := json.Unmarshal([]byte(horariosStr), &horariosInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Horarios inválidos"})
		return
	}

	if nombre == "" || descripcion == "" || categoria == "" || profesor == "" ||
		cupoTotalS == "" || len(horariosInput) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Faltan datos"})
		return
	}

	cupoTotalInt, err := strconv.Atoi(cupoTotalS)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cupo total inválido"})
		return
	}
	cupoTotal := uint(cupoTotalInt)

	var rutaImagen string
	if file, err := ctx.FormFile("imagen"); err == nil {
		filename := file.Filename
		dst := "../frontend/public/images/" + filename
		if err := ctx.SaveUploadedFile(file, dst); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar la imagen"})
			return
		}
		rutaImagen = "/images/" + filename
	}

	input := dto.ActividadConHorarioRequest{
		Nombre:      nombre,
		Descripcion: descripcion,
		Categoria:   categoria,
		Profesor:    profesor,
		Imagen:      rutaImagen,
		CupoTotal:   cupoTotal,
		Horarios:    horariosInput,
	}

	if err := services.CrearActividadConHorario(input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"msg": "Actividad creada"})
}

func EditarActividad(ctx *gin.Context) {
	rol, ok := ctx.Get("rol")
	if !ok || rol != "ADMIN" {
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

	// Validar que los datos requeridos estén presentes
	if input.Nombre == "" || input.Descripcion == "" || input.Categoria == "" || 
		input.Profesor == "" || input.CupoTotal == 0 || len(input.Horarios) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Faltan datos requeridos"})
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
