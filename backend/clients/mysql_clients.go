package clients

import (
	"backend/dao"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	user := "root"
	password := "Base041104"
	host := "localhost"
	port := 3306
	database := "backend"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		user, password, host, port, database)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("error connecting to DB: %v", err))
	}

	DB.AutoMigrate(&dao.User{})
	DB.Create(&dao.User{
		Id:           1,
		Username:     "emiliano",
		PasswordHash: "121j212hs9812sj2189sj",
	})

	DB.AutoMigrate(&dao.Actividad{})
	DB.Create(&dao.Actividad{
		Id:          1,
		Nombre:      "Actividad 1",
		Descripcion: "Descripcion de la actividad 1",
		Categoria:   "Categoria 1",
		CupoTotal:   10,
		Profesor:    "Profesor 1",
		Imagen:      "imagen1.jpg",
	})
	DB.AutoMigrate(&dao.Actividad{})
	DB.Create(&dao.Actividad{
		Id:          2,
		Nombre:      "Spinning",
		Descripcion: "Actividad en bici fija, guiada por profesor",
		Categoria:   "Categoria 1",
		CupoTotal:   20,
		Profesor:    "Francisca 1",
		Imagen:      "imagen1.jpg",
	})
	DB.AutoMigrate(&dao.Actividad{})
	DB.Create(&dao.Actividad{
		Id:          3,
		Nombre:      "Running",
		Descripcion: "Actividad al aire libre, guiada por profesor",
		Categoria:   "Categoria 2",
		CupoTotal:   20,
		Profesor:    "Magdalena Gomez",
		Imagen:      "imagen1.jpg",
	})
	DB.AutoMigrate(&dao.Horario{})
	DB.Create(&dao.Horario{
		Id:          1,
		IdActividad: 1,
		Dia:         "Lunes",
		HoraInicio:  time.Now(),
		HoraFin:     time.Now(),
		CupoHorario: nil,
	})

	DB.AutoMigrate(&dao.Inscripcion{})
	DB.Create(&dao.Inscripcion{
		Id:               1,
		IdUsuario:        1,
		IdHorario:        1,
		FechaInscripcion: time.Now(),
	})
}

func GetUserByUsername(username string) dao.User {
	var user dao.User
	// SELECT * FROM users WHERE username = ? LIMIT 1
	DB.First(&user, "username = ?", username)
	return user
}

func GetActividadById(id int) dao.Actividad {
	var actividad dao.Actividad
	// SELECT * FROM users WHERE username = ? LIMIT 1
	DB.First(&actividad, id)
	return actividad
}
