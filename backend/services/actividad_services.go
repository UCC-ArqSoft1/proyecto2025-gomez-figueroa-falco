package services

import (
	"backend/clients"
	"backend/dao"
	"backend/domain"

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

// InscribirUsuario crea un registro en tabla inscripciones.
func InscribirUsuario(userID, horarioID uint) error {
	ins := dao.Inscripcion{
		IdUsuario: userID,
		IdHorario: horarioID,
	}
	if err := clients.DB.Create(&ins).Error; err != nil {
		return err
	}
	return nil
}

func ActividadesDeUsuario(userID uint) ([]dao.Actividad, error) {
	var acts []dao.Actividad
	err := clients.DB.
		Joins("JOIN inscripciones ON inscripciones.IdHorario = horarios.Id").
		Joins("JOIN horarios       ON horarios.IdActividad = actividades.Id").
		Where("inscripciones.IdUsuario = ?", userID).
		Find(&acts)
	return acts, err.Error
}
