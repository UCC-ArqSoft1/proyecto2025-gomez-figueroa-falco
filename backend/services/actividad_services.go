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

	var act dao.Actividad
	clients.DB.Preload("Horarios").First(&act, id)

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
	var acts []dao.Actividad

	db := clients.DB.Model(&dao.Actividad{}).
		Preload("Horarios")

	if q != "" {
		like := "%" + q + "%"
		db = db.Joins("LEFT JOIN horarios h ON h.id_actividad = actividads.id").
			Where(`
                actividads.nombre      LIKE ?
			   OR actividads.profesor  LIKE ?
			   OR actividads.categoria LIKE ?
             OR DATE_FORMAT(h.hora_inicio, '%H:%i') LIKE ?`,
						like, like, like).
			Group("actividads.id") // evita duplicados
	}

	if err := db.Find(&acts).Error; err != nil {
		return nil, err
	}

	// Map a domain + horarios
	results := make([]domain.ActividadesDeportivas, 0, len(acts))
	for _, a := range acts {
		// armamos slice de horarios para la respuesta
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
			Horarios:    hs, // 👈 ahora el frontend los recibe
		})
	}

	return results, nil
}

func ActividadesDeUsuario(userID uint) ([]dao.Actividad, error) {
	var acts []dao.Actividad
	db := clients.DB.
		Preload("Horarios").
		Joins("JOIN horarios     ON horarios.id_actividad   = actividads.id").
		Joins("JOIN inscripcions ON inscripcions.id_horario = horarios.id").
		Where("inscripcions.id_usuario = ?", userID).
		Find(&acts)
	return acts, db.Error
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
