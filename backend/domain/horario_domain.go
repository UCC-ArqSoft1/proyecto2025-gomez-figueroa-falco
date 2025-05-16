package domain

import (
	"time"
)

type Horario struct {
	Id          uint      `gorm:"primaryKey;column:Id" json:"id"`
	Dia         string    `gorm:"type:enum('Lun','Mar','Mie','Jue','Vie','Sab','Dom');not null;column:Dia" json:"dia"`
	HoraInicio  time.Time `gorm:"type:time;not null;column:HoraInicio" json:"hora_inicio"`
	HoraFin     time.Time `gorm:"type:time;not null;column:HoraFin" json:"hora_fin"`
	IdActividad uint      `gorm:"not null;column:IdActividad" json:"id_actividad"`
	CupoHorario *uint     `gorm:"column:CupoHorario" json:"cupo_horario,omitempty"`

	Actividad Actividad `gorm:"foreignKey:IdActividad"`
}
