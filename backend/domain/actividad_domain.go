package domain

type ActividadesDeportivas struct {
	Id          uint   `json:"id" `
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion,omitempty"`
	Categoria   string `json:"categoria"`
	CupoTotal   uint   `json:"cupo_total"`
	Profesor    string `json:"profesor,omitempty"`
	Imagen      string `json:"imagen,omitempty"`

	// Relaciones
	Horarios      []Horario     ` json:"horarios"`
	Inscripciones []Inscripcion ` json:"inscripciones"`
}
