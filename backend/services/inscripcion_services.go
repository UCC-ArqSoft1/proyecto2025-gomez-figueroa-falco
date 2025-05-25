package services

import (
	"backend/clients"
	"backend/dao"
	"errors"
	"time"
)

// InscribirUsuario crea un registro en tabla inscripciones.
func InscribirUsuario(userID, horarioID uint) error {
	db := clients.DB

	// Verificar si el usuario existe
	var usuario dao.User
	if err := db.First(&usuario, userID).Error; err != nil {
		return errors.New("usuario no existe: " + err.Error())
	}

	// Verificar si el horario existe
	var horario dao.Horario
	if err := db.First(&horario, horarioID).Error; err != nil {
		return errors.New("horario no existe: " + err.Error())
	}

	//Verificar si el usuario ya está inscrito en el horario
	var inscrpcionExistente dao.Inscripcion
	if err := db.Where("id_usuario = ? AND id_horario = ?", userID, horarioID).First(&inscrpcionExistente).Error; err == nil {
		return errors.New("usuario ya inscrito en este horario")
	}

	// Verificar si hay cupo disponible
	if horario.CupoHorario != nil && *horario.CupoHorario <= 0 {
		return errors.New("cupo completo")
	}

	// Crear la inscripción
	// Si el cupo es mayor a 0, se crea la inscripción
	ins := dao.Inscripcion{
		IdUsuario:        userID,
		IdHorario:        horarioID,
		FechaInscripcion: time.Now(), // Asignar la fecha actual
	}
	if err := db.Create(&ins).Error; err != nil {
		return errors.New("error al crear inscripción: " + err.Error())
	}
	//Reducir en 1 el cupo disponible
	*horario.CupoHorario -= 1

	if err := db.Save(&horario).Error; err != nil {
		return errors.New("error al actualizar cupo horario: " + err.Error())
	}
	return nil
}
