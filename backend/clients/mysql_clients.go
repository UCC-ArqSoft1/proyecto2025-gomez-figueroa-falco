package clients

import (
	"backend/dao"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(fmt.Sprintf("Error convirtiendo DB_PORT: %v", err))
	}

	// Prints para debug
	fmt.Println("DB_USER:", user)
	fmt.Println("DB_PASS:", password)
	fmt.Println("DB_HOST:", host)
	fmt.Println("DB_PORT:", port)
	fmt.Println("DB_NAME:", database) 

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		user, password, host, port, database)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Error conectando a la base de datos: %v", err))
	}

	if err := DB.AutoMigrate(&dao.User{}); err != nil {
		panic(fmt.Sprintf("Error creando tabla: %s", err.Error()))
	}
	if err := DB.AutoMigrate(&dao.Horario{}); err != nil {
		panic(fmt.Sprintf("Error creando tabla: %s", err.Error()))
	}
	if err := DB.AutoMigrate(&dao.Actividad{}); err != nil {
		panic(fmt.Sprintf("Error creando tabla: %s", err.Error()))
	}
	if err := DB.AutoMigrate(&dao.Inscripcion{}); err != nil {
		panic(fmt.Sprintf("Error creando tabla: %s", err.Error()))
	}

	hashedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte("emiliano")))

	if result := DB.Where("username = ?", "emiliano").FirstOrCreate(&dao.User{
		Nombre:       "Emiliano",
		Username:     "emiliano",
		Email:        "emiliano@gmial.com",
		PasswordHash: hashedPassword,
		Rol:          "SOCIO",
	}); result.Error != nil {
		fmt.Println("Error creando usuario: ", result.Error)
	}

	contraseña := fmt.Sprintf("%x", sha256.Sum256([]byte("fran")))

	if result := DB.Where("username = ?", "francisca").FirstOrCreate(&dao.User{
		Nombre:       "Francisca",
		Username:     "francisca",
		Email:        "franciscafalco4@gmail.com",
		PasswordHash: contraseña,
		Rol:          "ADMIN",
	}); result.Error != nil {
		fmt.Println("Error creando usuario: ", result.Error)
	}
	cupo10 := uint(10)
	if result := DB.Where("nombre = ?", "Funcional").FirstOrCreate(&dao.Actividad{
		Nombre:      "Funcional",
		Descripcion: "Entrenaminento basado en movimeintos naturales del cuerpo. Mejora tu eficiencia diaria aprendiendo a moverte mejor y prevenir lesiones",
		Categoria:   "Categoria 1",
		CupoTotal:   10,
		Profesor:    "Juan Cabral",
		Imagen:      "/images/funcional.png",
		Horarios: []dao.Horario{
			{
				Dia:         "Lunes",
				HoraInicio: time.Date(2025, 7, 7, 18, 0, 0, 0, time.Local),
				HoraFin:    time.Date(2025, 7, 7, 19, 0, 0, 0, time.Local),
				CupoHorario: &cupo10,
			},
			{
				Dia:         "Martes",
				HoraInicio:  time.Date(2025, 7, 8, 18, 0, 0, 0, time.Local),
				HoraFin:    time.Date(2025, 7, 8, 19, 0, 0, 0, time.Local),
				CupoHorario: &cupo10,
			},
		},
	}); result.Error != nil {
		fmt.Println("Error creando actividad: ", result.Error)
	}
 cupo20 := uint(20)
	if result := DB.Where("nombre = ?", "Pilates").FirstOrCreate(&dao.Actividad{
		Nombre:      "Pilates",
		Descripcion: "Es un método de entrenamiento que utiliza el propio peso corporal para fortalecer y tonificar los músculos, mejorar la postura y la flexibilidad, y aumentar la resistencia física y mental.",
		Categoria:   "Categoria 1",
		CupoTotal:   20,
		Profesor:    "Francisca 1",
		Imagen:      "/images/pilates.png",
		Horarios: []dao.Horario{
			{
				Dia:         "Martes",
				HoraInicio:  time.Date(2025, 7, 8, 15, 0, 0, 0, time.Local),
				HoraFin:    time.Date(2025, 7, 8, 16, 0, 0, 0, time.Local),
				CupoHorario: &cupo20,
			},
			{
				Dia:         "Jueves",
				HoraInicio:  time.Date(2025, 7, 10, 17, 0, 0, 0, time.Local),
				HoraFin:   	time.Date(2025, 7, 10, 18, 0, 0, 0, time.Local),
				CupoHorario: &cupo20,
			},
		},
	}); result.Error != nil {
		fmt.Println("Error creando actividad: ", result.Error)
	}

	if result := DB.Where("nombre = ?", "Spinning").FirstOrCreate(&dao.Actividad{
		Nombre:      "Spinning",
		Descripcion: "Clase de ciclismo indoor con música motivadora. Quema calorías, mejora tu resistencia cardiovascular y tonifica tus piernas en un ambiente energético y divertido.",
		Categoria:   "Categoria 2",
		CupoTotal:   20,
		Profesor:    "Magdalena Gomez",
		Imagen:      "/images/spining.png",
		Horarios: []dao.Horario{
			{
				Dia:         "Lunes",
				HoraInicio:  time.Date(2025, 7, 7, 9, 0, 0, 0, time.Local),
				HoraFin:     time.Date(2025, 7, 7, 10, 0, 0, 0, time.Local),
				CupoHorario: &cupo20,
			},
			{
				Dia:         "Miercoles",
				HoraInicio:  time.Date(2025, 7, 9, 9, 0, 0, 0, time.Local),
				HoraFin:     time.Date(2025, 7, 9, 10, 0, 0, 0, time.Local),
				CupoHorario: &cupo20,
			},
		},
	}); result.Error != nil {
		fmt.Println("Error creando actividad: ", result.Error)
	}

}

