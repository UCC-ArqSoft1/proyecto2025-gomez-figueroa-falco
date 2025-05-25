package clients

import (
	"backend/dao"
	"crypto/sha256"
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

	if err := DB.AutoMigrate(&dao.User{}); err != nil {
		panic(fmt.Sprintf("Error creating table: %s", err.Error()))
	}
	if err := DB.AutoMigrate(&dao.Horario{}); err != nil {
		panic(fmt.Sprintf("Error creating table: %s", err.Error()))
	}
	if err := DB.AutoMigrate(&dao.Actividad{}); err != nil {
		panic(fmt.Sprintf("Error creating table: %s", err.Error()))
	}
	if err := DB.AutoMigrate(&dao.Inscripcion{}); err != nil {
		panic(fmt.Sprintf("Error creating table: %s", err.Error()))
	}

	hashedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte("admin")))

	if result := DB.Create(&dao.User{
		Username:     "emiliano",
		PasswordHash: hashedPassword,
	}); result.Error != nil {
		fmt.Println("Error creating user: ", result.Error)
	}

	if result := DB.Create(&dao.Actividad{
		Nombre:      "Funcional",
		Descripcion: "Descripcion de la actividad 1",
		Categoria:   "Categoria 1",
		CupoTotal:   10,
		Profesor:    "Juan Cabral",
		Imagen:      "imagen1.jpg",
		Horarios: []dao.Horario{
			{
				Dia:         "Lunes",
				HoraInicio:  time.Now(),
				HoraFin:     time.Now(),
				CupoHorario: nil,
			},
			{
				Dia:         "Martes",
				HoraInicio:  time.Now(),
				HoraFin:     time.Now(),
				CupoHorario: nil,
			},
		},
	}); result.Error != nil {
		fmt.Println("Error creating activity: ", result.Error)
	}

	if result := DB.Create(&dao.Actividad{
		Nombre:      "Spinning",
		Descripcion: "Actividad en bici fija, guiada por profesor",
		Categoria:   "Categoria 1",
		CupoTotal:   20,
		Profesor:    "Francisca 1",
		Imagen:      "imagen1.jpg",
		Horarios: []dao.Horario{
			{
				Dia:         "Martes",
				HoraInicio:  time.Now(),
				HoraFin:     time.Now(),
				CupoHorario: nil,
			},
			{
				Dia:         "Jueves",
				HoraInicio:  time.Now(),
				HoraFin:     time.Now(),
				CupoHorario: nil,
			},
		},
	}); result.Error != nil {
		fmt.Println("Error creating activity: ", result.Error)
	}

	if result := DB.Create(&dao.Actividad{
		Nombre:      "Running",
		Descripcion: "Actividad al aire libre, guiada por profesor",
		Categoria:   "Categoria 2",
		CupoTotal:   20,
		Profesor:    "Magdalena Gomez",
		Imagen:      "imagen1.jpg",
		Horarios: []dao.Horario{
			{
				Dia:         "Lunes",
				HoraInicio:  time.Now(),
				HoraFin:     time.Now(),
				CupoHorario: nil,
			},
			{
				Dia:         "Miercoles",
				HoraInicio:  time.Now(),
				HoraFin:     time.Now(),
				CupoHorario: nil,
			},
		},
	}); result.Error != nil {
		fmt.Println("Error creating activity: ", result.Error)
	}

	DB.Create(&dao.Inscripcion{
		Dia:              "Lunes",
		HoraInicio:       time.Now(),
		HoraFin:          time.Now(),
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
