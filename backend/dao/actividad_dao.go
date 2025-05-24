package dao

type Actividad struct {
	Id          uint   `gorm:"primaryKey:autoIncrement;column:id"`
	Nombre      string `gorm:"size:100;not null;column:nombre"`
	Descripcion string `gorm:"type:text;column:descripcion"`
	Categoria   string `gorm:"size:50;not null;column:categoria"`
	CupoTotal   uint   `gorm:"not null;column:cupo_total"`
	Profesor    string `gorm:"size:100;column:profesor"`
	Imagen      string `gorm:"size:255;column:imagen"`

	Horarios []Horario `gorm:"foreignKey:IdActividad;references:Id;constraint:OnDelete:CASCADE"`
}
