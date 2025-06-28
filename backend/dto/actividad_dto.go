package dto

type HorarioRequest struct {
	Dia         string `json:"dia"`
	HoraInicio  string `json:"hora_inicio"`
	HoraFin     string `json:"hora_fin"`
	CupoHorario *uint  `json:"cupo_horario"`
}

type ActividadConHorarioRequest struct {
	Nombre      string           `json:"nombre"`
	Descripcion string           `json:"descripcion"`
	Categoria   string           `json:"categoria"`
	Profesor    string           `json:"profesor"`
	Imagen      string           `json:"imagen"`
	CupoTotal   uint             `json:"cupo_total"`
	Horarios    []HorarioRequest `json:"horarios"`
}
