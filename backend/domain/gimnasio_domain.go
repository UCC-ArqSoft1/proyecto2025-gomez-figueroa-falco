package domain

type ActividadesDeportivas struct {
	Horarios  Horarios `json:"horario"`
	Cupo      int      `json:"cupo"`
	Categoria string   `json:"categoria"`
}

type Horarios struct {
	Dias []DiaSemana
	Hora []Horas
}

type DiaSemana string

const (
	Lunes     DiaSemana = "Lunes"
	Martes    DiaSemana = "Martes"
	Miercoles DiaSemana = "Miercoles"
	Jueves    DiaSemana = "Jueves"
	Viernes   DiaSemana = "Viernes"
	Sabado    DiaSemana = "Sabado"
)

type Horas struct {
	Empieza string `json:"empieza"`
	Termina string `json:"termina"`
}
