package dao

import (
	"time"
	"gorm.io/gorm"
)

type Horario struct {
	Id          uint      `gorm:"primaryKey:autoIncrement;column:id"`
	Dia         string    `gorm:"not null;column:dia"`
	HoraInicio  time.Time `gorm:"type:time;not null;column:hora_inicio"`
	HoraFin     time.Time `gorm:"type:time;not null;column:hora_fin"`
	IdActividad uint      `gorm:"not null;column:id_actividad"`
	CupoHorario *uint     `gorm:"column:cupo_horario"`
}

func CrearHorario(db *gorm.DB, horario Horario) error {
	return db.Create(&horario).Error
}
