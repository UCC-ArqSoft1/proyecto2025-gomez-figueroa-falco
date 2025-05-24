package domain

import (
	"time"
)

type Horario struct {
	Id          uint      `json:"id"`
	Dia         string    `json:"dia"`
	HoraInicio  time.Time `json:"hora_inicio"`
	HoraFin     time.Time `json:"hora_fin"`
	IdActividad uint      `json:"id_actividad"`
	CupoHorario *uint     `json:"cupo_horario,omitempty"`
}
