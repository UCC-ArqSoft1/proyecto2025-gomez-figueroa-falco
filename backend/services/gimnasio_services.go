package services

import (
	"backend/domain"
)

func GetActividadById(id int) domain.ActividadesDeportivas {

	horarios := domain.Horarios{
		Dias: []domain.DiaSemana{domain.Lunes, domain.Martes},
		Hora: []domain.Horas{
			{
				Empieza: "18:00",
				Termina: "19:00",
			},
			{
				Empieza: "21:00",
				Termina: "22:00",
			},
		},
	}

	actividad := domain.ActividadesDeportivas{
		Horarios:  horarios,
		Cupo:      20,
		Categoria: "Sppining",
	}
	return actividad
}
