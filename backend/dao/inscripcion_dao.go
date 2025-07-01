package dao

import (
	"time"
	"gorm.io/gorm"
)

type Inscripcion struct {
	Id               uint      `gorm:"primaryKey:autIncrement;column:id"`
	Dia              string    `gorm:"not null;column:dia"`
	HoraInicio       time.Time `gorm:"type:time;not null;column:hora_inicio"`
	HoraFin          time.Time `gorm:"type:time;not null;column:hora_fin"`
	FechaInscripcion time.Time `gorm:"column:fecha_inscripcion;autoCreateTime"`

	IdUsuario   uint `gorm:"not null;column:id_usuario"`
	IdHorario   uint `gorm:"not null;column:id_horario"`
	IdActividad uint `gorm:"not null;column:id_actividad"`
}

func BuscarUsuarioPorID(db *gorm.DB, userID uint) (User, error) {
	var usuario User
	err := db.First(&usuario, userID).Error
	return usuario, err
}

func BuscarHorarioPorID(db *gorm.DB, horarioID uint) (Horario, error) {
	var horario Horario
	err := db.First(&horario, horarioID).Error
	return horario, err
}

func BuscarInscripcionExistente(db *gorm.DB, userID, horarioID uint) (Inscripcion, error) {
	var inscripcion Inscripcion
	err := db.Where("id_usuario = ? AND id_horario = ?", userID, horarioID).First(&inscripcion).Error
	return inscripcion, err
}

func CrearInscripcion(db *gorm.DB, ins Inscripcion) error {
	return db.Create(&ins).Error
}

func ActualizarHorario(db *gorm.DB, horario *Horario) error {
	return db.Save(horario).Error
}
