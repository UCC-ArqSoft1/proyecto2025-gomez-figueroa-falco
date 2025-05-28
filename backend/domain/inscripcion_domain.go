package domain

import (
	"time"
)

type Inscripcion struct {
	Id               uint      `json:"id"`
	Dia              string    `json:"dia"`
	HoraInicio       time.Time `json:"hora_inicio"`
	HoraFin          time.Time `json:"hora_fin"`
	IdUsuario        uint      `json:"id_usuario"`
	IdHorario        uint      `json:"id_horario"`
	IdActividad      uint      `json:"id_actividad"`
	FechaInscripcion time.Time `json:"fecha_inscripcion"`

	// Relaciones
	Usuario User
	Horario Horario
}
