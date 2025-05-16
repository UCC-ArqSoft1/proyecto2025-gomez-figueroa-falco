package services

import (
	"backend/clients"
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
