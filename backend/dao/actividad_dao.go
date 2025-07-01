package dao

import (
	"gorm.io/gorm"
)

type Actividad struct {
	Id          uint   `gorm:"primaryKey:autoIncrement;column:id"`
	Nombre      string `gorm:"size:100;not null;column:nombre"`
	Descripcion string `gorm:"type:text;column:descripcion"`
	Categoria   string `gorm:"size:50;not null;column:categoria"`
	CupoTotal   uint   `gorm:"not null;column:cupo_total"`
	Profesor    string `gorm:"size:100;column:profesor"`
	Imagen      string `gorm:"size:255;column:imagen"`

	Horarios []Horario `gorm:"foreignKey:IdActividad;references:Id;constraint:OnDelete:CASCADE"`
}

func GetActividadById(db *gorm.DB, id int) (Actividad, error) {
	var act Actividad
	err := db.Preload("Horarios").First(&act, id).Error
	return act, err
}

func BuscarActividades(db *gorm.DB, q string) ([]Actividad, error) {
	var acts []Actividad
	dbq := db.Model(&Actividad{}).Preload("Horarios")
	if q != "" {
		like := "%" + q + "%"
		dbq = dbq.Joins("LEFT JOIN horarios h ON h.id_actividad = actividads.id").
			Where(`actividads.nombre LIKE ? OR actividads.profesor LIKE ? OR DATE_FORMAT(h.hora_inicio, '%H:%i') LIKE ?`, like, like, like).
			Group("actividads.id")
	}
	if err := dbq.Find(&acts).Error; err != nil {
		return nil, err
	}
	return acts, nil
}

func BuscarActividadesDeUsuario(db *gorm.DB, userID uint) ([]Actividad, error) {
	var acts []Actividad
	dbq := db.Preload("Horarios", "id IN (SELECT id_horario FROM inscripcions WHERE id_usuario = ?)", userID).
		Joins("JOIN inscripcions ON inscripcions.id_actividad = actividads.id").
		Where("inscripcions.id_usuario = ?", userID).
		Group("actividads.id")
	dbq.Find(&acts)
	return acts, dbq.Error
}

func CrearActividad(db *gorm.DB, act *Actividad) error {
	return db.Create(act).Error
}

func ActualizarActividad(db *gorm.DB, act *Actividad) error {
	return db.Save(act).Error
}

func EliminarActividad(db *gorm.DB, id uint) error {
	return db.Delete(&Actividad{}, id).Error
}

func EliminarHorariosPorActividad(db *gorm.DB, actividadID uint) error {
	return db.Where("id_actividad = ?", actividadID).Delete(&Horario{}).Error
}
