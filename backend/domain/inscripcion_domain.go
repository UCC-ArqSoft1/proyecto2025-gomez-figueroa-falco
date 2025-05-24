package domain

import (
	"time"
)

type Inscripcion struct {
	Id               uint      `json:"id"`
	IdUsuario        uint      `json:"id_usuario"`
	IdHorario        uint      `json:"id_horario"`
	FechaInscripcion time.Time `json:"fecha_inscripcion"`

	// Relaciones
	Usuario User
	Horario Horario
}
