package services

import (
	"backend/clients"
	"backend/dao"
	"backend/domain"
	"backend/dto"
	"errors"
	"time"

	"gorm.io/gorm"
)

func GetActividadById(id int) domain.ActividadesDeportivas {

	actDAO := clients.GetActividadById(id)

	return domain.ActividadesDeportivas{
		Id:          actDAO.Id,
		Nombre:      actDAO.Nombre,
		Descripcion: actDAO.Descripcion,
		Categoria:   actDAO.Categoria,
		CupoTotal:   actDAO.CupoTotal,
		Profesor:    actDAO.Profesor,
		Imagen:      actDAO.Imagen,
	}
}

// filtra por palabra clave en título o categoría
// devuelve todas las actividades si no se pasa palabra clave
func BuscarActividades(q string) ([]domain.ActividadesDeportivas, error) {
	var acts []dao.Actividad
	db := clients.DB.Preload("Horarios")
	results := make([]domain.ActividadesDeportivas, 0)
	var err *gorm.DB
	if q == "" {
		// trae todas las actividades
		err = db.Find(&acts)
	} else {
		pattern := "%" + q + "%"
		err = db.Where("Nombre LIKE ? OR categoria LIKE ?", pattern, pattern).Find(&acts)
	}
	if err.Error != nil {
		return nil, err.Error
	}

	for _, actDAO := range acts {
		results = append(results, domain.ActividadesDeportivas{
			Id:          actDAO.Id,
			Nombre:      actDAO.Nombre,
			Descripcion: actDAO.Descripcion,
			Categoria:   actDAO.Categoria,
			CupoTotal:   actDAO.CupoTotal,
			Profesor:    actDAO.Profesor,
			Imagen:      actDAO.Imagen,
		})
	}
	return results, nil
}

func ActividadesDeUsuario(userID uint) ([]dao.Actividad, error) {
	var acts []dao.Actividad
	err := clients.DB.
		Joins("JOIN horarios       ON horarios.IdActividad = actividades.Id").
		Joins("JOIN inscripciones ON inscripciones.IdHorario = horarios.Id").
		Where("inscripciones.IdUsuario = ?", userID).
		Find(&acts)
	return acts, err.Error
}

func CrearActividadConHorario(input dto.ActividadConHorarioRequest) error {
	db := clients.DB

	return db.Transaction(func(tx *gorm.DB) error {
		// Crear la actividad
		actividad := dao.Actividad{
			Nombre:      input.Nombre,
			Descripcion: input.Descripcion,
			Categoria:   input.Categoria,
			Profesor:    input.Profesor,
			Imagen:      input.Imagen,
			CupoTotal:   input.CupoTotal,
		}
		if err := tx.Create(&actividad).Error; err != nil {
			return err
		}

		//Parsear hora de inicio y fin
		layout := "15:04"
		horaInicio, err := time.Parse(layout, input.Horario.HoraInicio)
		if err != nil {
			return errors.New("error al parsear hora de inicio: " + err.Error())
		}
		horaFin, err := time.Parse(layout, input.Horario.HoraFin)
		if err != nil {
			return errors.New("error al parsear hora de fin: " + err.Error())
		}

		// Crear el horario
		horario := dao.Horario{
			Dia:         input.Horario.Dia,
			HoraInicio:  horaInicio,
			HoraFin:     horaFin,
			CupoHorario: input.Horario.CupoHorario,
			IdActividad: actividad.Id,
		}
		if err := tx.Create(&horario).Error; err != nil {
			return err
		}
		return nil
	})
}

func ActualizarActividad(id uint, input dto.ActividadConHorarioRequest) error {
	db := clients.DB

	// Buscar la actividad por ID
	var actividad dao.Actividad
	if err := db.First(&actividad, id).Error; err != nil {
		return errors.New("actividad no encontrada")
	}

	// Actualizar los campos de la actividad
	actividad.Nombre = input.Nombre
	actividad.Descripcion = input.Descripcion
	actividad.Categoria = input.Categoria
	actividad.Profesor = input.Profesor
	actividad.Imagen = input.Imagen
	actividad.CupoTotal = input.CupoTotal

	return db.Save(&actividad).Error
}
func EliminarActividad(id uint) error {
	db := clients.DB
	if err := db.Delete(&dao.Actividad{}, id).Error; err != nil {
		return errors.New("error al eliminar actividad: " + err.Error())
	}
	return nil
}
