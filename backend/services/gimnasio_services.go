package services

import (
	"backend/domain"
	"time"
)

func GetActividadById(id int) domain.ActividadesDeportivas {

	start1, _ := time.Parse("15:04", "18:00")
	end1, _ := time.Parse("15:04", "19:00")
	parte1 := domain.Horarios{
		Dias: []domain.DiaSemana{domain.Lunes},
		Hora: domain.Horas{
			Empieza: start1,
			Termina: end1,
		},
	}

	actividad := domain.ActividadesDeportivas{
		Horarios:  []domain.Horarios{parte1},
		Cupo:      20,
		Categoria: "Sppining",
	}
	return actividad
}
