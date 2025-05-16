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
