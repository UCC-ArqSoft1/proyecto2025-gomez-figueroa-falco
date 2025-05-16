package dao

import (
	"time"
)

type Inscripcion struct {
	Id               uint      `gorm:"primaryKey;column:Id" json:"id"`
	IdUsuario        uint      `gorm:"not null;column:IdUsuario" json:"id_usuario"`
	IdHorario        uint      `gorm:"not null;column:IdHorario" json:"id_horario"`
	FechaInscripcion time.Time `gorm:"column:FechaInscripcion;autoCreateTime" json:"fecha_inscripcion"`
}
