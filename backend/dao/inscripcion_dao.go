package dao

import (
	"time"
)

type Inscripcion struct {
	Id               uint      `gorm:"primaryKey:autIncrement;column:id"`
	Dia              string    `gorm:"not null;column:dia"`
	HoraInicio       time.Time `gorm:"type:time;not null;column:hora_inicio"`
	HoraFin          time.Time `gorm:"type:time;not null;column:hora_fin"`
	FechaInscripcion time.Time `gorm:"column:fecha_inscripcion;autoCreateTime"`
}
