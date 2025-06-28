package services

import (
	"backend/clients"
	"backend/dao"
	"backend/domain"
	"backend/dto"
	"errors"
	"time"
"fmt"
	"gorm.io/gorm"
)

func GetActividadById(id int) domain.ActividadesDeportivas {
	var act dao.Actividad
	clients.DB.Preload("Horarios").First(&act, id)

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
	var acts []dao.Actividad

	db := clients.DB.Model(&dao.Actividad{}).
		Preload("Horarios")

	if q != "" {
		like := "%" + q + "%"
		db = db.Joins("LEFT JOIN horarios h ON h.id_actividad = actividads.id").
			Where(`
                actividads.nombre      LIKE ?
			   OR actividads.profesor  LIKE ?
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
		Preload("Horarios", "id IN (SELECT id_horario FROM inscripcions WHERE id_usuario = ?)", userID).
		Joins("JOIN inscripcions ON inscripcions.id_actividad = actividads.id").
		Where("inscripcions.id_usuario = ?", userID).
		Group("actividads.id").
		Find(&acts)
	return acts, db.Error
}

func CrearActividadConHorario(input dto.ActividadConHorarioRequest) error {
	db := clients.DB

	// Array de días en español
	diasSemana := []string{"Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"}

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

		
		// Guardar todos los horarios
		for _, h := range input.Horarios {

			fmt.Printf("Parseando hora_inicio: '%s', hora_fin: '%s'\n", h.HoraInicio, h.HoraFin)
				// ... resto del código ...
			
			layout := "2006-01-02 15:04"
			
			
			horaInicio, err := time.Parse(layout, h.HoraInicio)
			if err != nil {
				return errors.New("error al parsear hora de inicio: " + err.Error())
			}
			horaFin, err := time.Parse(layout, h.HoraFin)
			if err != nil {
				return errors.New("error al parsear hora de fin: " + err.Error())
			}
			dia := diasSemana[int(horaInicio.Weekday())]
			horario := dao.Horario{
				Dia:         dia,
				HoraInicio:  horaInicio,
				HoraFin:     horaFin,
				CupoHorario: h.CupoHorario,
				IdActividad: actividad.Id,
			}
			if err := tx.Create(&horario).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func ActualizarActividad(id uint, input dto.ActividadConHorarioRequest) error {
	db := clients.DB

	// Array de días en español
	diasSemana := []string{"Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"}

	return db.Transaction(func(tx *gorm.DB) error {
		// Buscar la actividad por ID
		var actividad dao.Actividad
		if err := tx.First(&actividad, id).Error; err != nil {
			return errors.New("actividad no encontrada")
		}

		// Actualizar los campos de la actividad
		actividad.Nombre = input.Nombre
		actividad.Descripcion = input.Descripcion
		actividad.Categoria = input.Categoria
		actividad.Profesor = input.Profesor
		actividad.Imagen = input.Imagen
		actividad.CupoTotal = input.CupoTotal

		if err := tx.Save(&actividad).Error; err != nil {
			return err
		}

		// Eliminar todos los horarios existentes de la actividad
		if err := tx.Where("id_actividad = ?", id).Delete(&dao.Horario{}).Error; err != nil {
			return err
		}

		// Crear los nuevos horarios
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
			
			dia := diasSemana[int(horaInicio.Weekday())]
			horario := dao.Horario{
				Dia:         dia,
				HoraInicio:  horaInicio,
				HoraFin:     horaFin,
				CupoHorario: h.CupoHorario,
				IdActividad: actividad.Id,
			}
			if err := tx.Create(&horario).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
func EliminarActividad(id uint) error {
	db := clients.DB
	if err := db.Delete(&dao.Actividad{}, id).Error; err != nil {
		return errors.New("error al eliminar actividad: " + err.Error())
	}
	return nil
}
