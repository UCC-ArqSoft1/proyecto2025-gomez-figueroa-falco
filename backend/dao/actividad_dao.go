package dao

type Actividad struct {
	Id          uint   `gorm:"primaryKey;column:Id" json:"id" `
	Nombre      string `gorm:"size:100;not null;column:Nombre" json:"nombre"`
	Descripcion string `gorm:"type:text;column:Descripcion" json:"descripcion,omitempty"`
	Categoria   string `gorm:"size:50;not null;column:Categoria" json:"categoria"`
	CupoTotal   uint   `gorm:"not null;column:CupoTotal" json:"cupo_total"`
	Profesor    string `gorm:"size:100;column:Profesor" json:"profesor,omitempty"`
	Imagen      string `gorm:"size:255;column:Imagen"  json:"imagen,omitempty"`

	Horarios []Horario `gorm:"foreignKey:IdActividad;references:Id" json:"horarios"`
}
