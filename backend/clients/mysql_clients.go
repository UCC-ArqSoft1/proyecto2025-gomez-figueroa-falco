package clients

import (
	"backend/dao"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	user := "root"
	password := "1234"
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
	DB.AutoMigrate(&dao.Horario{})
	DB.Create(&dao.Horario{
		Id:          1,
		IdActividad: 1,
		Dia:         "Lunes",
		HoraInicio:  "10:00",
		HoraFin:     "11:00",
		CupoHorario: nil,
	})

	DB.AutoMigrate(&dao.Inscripcion{})
	DB.Create(&dao.Inscripcion{
		Id:               1,
		IdUsuario:        1,
		IdHorario:        1,
		FechaInscripcion: "2023-10-01",
	})
}

func GetUserByUsername(username string) dao.User {
	var user dao.User
	// SELECT * FROM users WHERE username = ? LIMIT 1
	DB.First(&user, "username = ?", username)
	return user
}
