package services

import (
	"backend/clients"
	"backend/dao"
	horarioDao "backend/dao"
	"backend/domain"
	"backend/dto"
	"errors"
	"time"
	"gorm.io/gorm"
)

func GetActividadById(id int) domain.ActividadesDeportivas {
	act, err := dao.GetActividadById(clients.DB, id)
	if err != nil {
		// Manejo simple de error: devolver vacío (puedes mejorarlo)
		return domain.ActividadesDeportivas{}
	}

	// Convertir todos los horarios al dominio
	horarios := make([]domain.Horario, len(act.Horarios))
	for i, h := range act.Horarios {
		horarios[i] = domain.Horario{
			Id:          h.Id,
			Dia:         h.Dia,
			HoraInicio:  h.HoraInicio,
			HoraFin:     h.HoraFin,
			IdActividad: h.IdActividad,
			CupoHorario: h.CupoHorario,
		}
	}

	return domain.ActividadesDeportivas{
		Id:          act.Id,
		Nombre:      act.Nombre,
		Descripcion: act.Descripcion,
		Categoria:   act.Categoria,
		CupoTotal:   act.CupoTotal,
		Profesor:    act.Profesor,
		Imagen:      act.Imagen,
		Horarios:    horarios,
	}
}

// BuscarActividades filtra por nombre, categoría o HORA (HH:mm)
// Si q == "", devuelve todas
func BuscarActividades(q string) ([]domain.ActividadesDeportivas, error) {
	acts, err := dao.BuscarActividades(clients.DB, q)
	if err != nil {
		return nil, err
	}
	results := make([]domain.ActividadesDeportivas, 0, len(acts))
	for _, a := range acts {
		hs := make([]domain.Horario, 0, len(a.Horarios))
		for _, h := range a.Horarios {
			hs = append(hs, domain.Horario{
				Id:         h.Id,
				Dia:        h.Dia,
				HoraInicio: h.HoraInicio,
				HoraFin:    h.HoraFin,
			})
		}
		results = append(results, domain.ActividadesDeportivas{
			Id:          a.Id,
			Nombre:      a.Nombre,
			Descripcion: a.Descripcion,
			Categoria:   a.Categoria,
			CupoTotal:   a.CupoTotal,
			Profesor:    a.Profesor,
			Imagen:      a.Imagen,
			Horarios:    hs,
		})
	}
	return results, nil
}

func ActividadesDeUsuario(userID uint) ([]dao.Actividad, error) {
	return dao.BuscarActividadesDeUsuario(clients.DB, userID)
}

func CrearActividadConHorario(input dto.ActividadConHorarioRequest) error {
	db := clients.DB
	return db.Transaction(func(tx *gorm.DB) error {
		actividad := dao.Actividad{
			Nombre:      input.Nombre,
			Descripcion: input.Descripcion,
			Categoria:   input.Categoria,
			Profesor:    input.Profesor,
			Imagen:      input.Imagen,
			CupoTotal:   input.CupoTotal,
		}
		if err := dao.CrearActividad(tx, &actividad); err != nil {
			return err
		}
		// Guardar todos los horarios
		for _, h := range input.Horarios {
			layout := "2006-01-02 15:04"
			horaInicio, err := time.Parse(layout, h.HoraInicio)
			if err != nil {
				return errors.New("error al parsear hora de inicio: " + err.Error())
			}
			horaFin, err := time.Parse(layout, h.HoraFin)
			if err != nil {
				return errors.New("error al parsear hora de fin: " + err.Error())
			}
			dia := []string{"Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"}[int(horaInicio.Weekday())]
			horario := horarioDao.Horario{
				Dia:         dia,
				HoraInicio:  horaInicio,
				HoraFin:     horaFin,
				CupoHorario: h.CupoHorario,
				IdActividad: actividad.Id,
			}
			if err := horarioDao.CrearHorario(tx, horario); err != nil {
				return err
			}
		}
		return nil
	})
}

func ActualizarActividad(id uint, input dto.ActividadConHorarioRequest) error {
	db := clients.DB
	return db.Transaction(func(tx *gorm.DB) error {
		var actividad dao.Actividad
		if err := tx.First(&actividad, id).Error; err != nil {
			return errors.New("actividad no encontrada")
		}
		actividad.Nombre = input.Nombre
		actividad.Descripcion = input.Descripcion
		actividad.Categoria = input.Categoria
		actividad.Profesor = input.Profesor
		actividad.Imagen = input.Imagen
		actividad.CupoTotal = input.CupoTotal
		if err := dao.ActualizarActividad(tx, &actividad); err != nil {
			return err
		}
		if err := dao.EliminarHorariosPorActividad(tx, id); err != nil {
			return err
		}
		for _, h := range input.Horarios {
			layout := "2006-01-02 15:04"
			horaInicio, err := time.Parse(layout, h.HoraInicio)
			if err != nil {
				return errors.New("error al parsear hora de inicio: " + err.Error())
			}
			horaFin, err := time.Parse(layout, h.HoraFin)
			if err != nil {
				return errors.New("error al parsear hora de fin: " + err.Error())
			}
			dia := []string{"Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"}[int(horaInicio.Weekday())]
			horario := horarioDao.Horario{
				Dia:         dia,
				HoraInicio:  horaInicio,
				HoraFin:     horaFin,
				CupoHorario: h.CupoHorario,
				IdActividad: actividad.Id,
			}
			if err := horarioDao.CrearHorario(tx, horario); err != nil {
				return err
			}
		}
		return nil
	})
}

func EliminarActividad(id uint) error {
	return dao.EliminarActividad(clients.DB, id)
}
