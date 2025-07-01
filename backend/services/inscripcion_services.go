package services

import (
	"backend/clients"
	"backend/dao"
	"errors"
	"time"
)

// InscribirUsuario crea un registro en tabla inscripciones.
func InscribirUsuario(userID, horarioID, actividadID uint, dia string) error {
	// Verificar si el usuario existe
	_, err := dao.BuscarUsuarioPorID(clients.DB, userID)
	if err != nil {
		return errors.New("usuario no existe: " + err.Error())
	}

	// Verificar si el horario existe
	horario, err := dao.BuscarHorarioPorID(clients.DB, horarioID)
	if err != nil {
		return errors.New("horario no existe: " + err.Error())
	}

	// Verificar si el horario tiene cupo inicializado
	if horario.CupoHorario == nil {
		return errors.New("El horario no tiene cupo inicializado. Contacte al administrador.")
	}

	//Verificar si el usuario ya está inscrito en el horario
	_, err = dao.BuscarInscripcionExistente(clients.DB, userID, horarioID)
	if err == nil {
		return errors.New("usuario ya inscrito en este horario")
	}

	// Verificar si hay cupo disponible
	if horario.CupoHorario != nil && *horario.CupoHorario <= 0 {
		return errors.New("cupo completo")
	}

	// Crear la inscripción
	ins := dao.Inscripcion{
		Dia:              dia,
		HoraInicio:       horario.HoraInicio,
		HoraFin:          horario.HoraFin,
		IdUsuario:        userID,
		IdHorario:        horarioID,
		IdActividad:      actividadID,
		FechaInscripcion: time.Now(),
	}
	if err := dao.CrearInscripcion(clients.DB, ins); err != nil {
		return errors.New("error al crear inscripción: " + err.Error())
	}
	//Reducir en 1 el cupo disponible
	*horario.CupoHorario--

	if err := dao.ActualizarHorario(clients.DB, &horario); err != nil {
		return errors.New("error al actualizar cupo horario: " + err.Error())
	}
	return nil
}
