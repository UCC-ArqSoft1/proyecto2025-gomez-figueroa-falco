package services

import (
	"backend/clients"
	"backend/dao"
	"backend/domain"
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
func BuscarActividades(q string) ([]dao.Actividad, error) {
	var acts []dao.Actividad
	db := clients.DB.Preload("Horarios")
	if q == "" {
		return acts, db.Find(&acts).Error
	}
	pattern := "%" + q + "%"
	return acts, db.Where("Nombre LIKE ? OR categoria LIKE ?", pattern, pattern).Find(&acts).Error
}

// InscribirUsuario crea un registro en tabla inscripciones.
func InscribirUsuario(userID, horarioID uint) error {
	ins := dao.Inscripcion{IdUsuario: userID, IdHorario: horarioID}
	return clients.DB.Create(&ins).Error
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
