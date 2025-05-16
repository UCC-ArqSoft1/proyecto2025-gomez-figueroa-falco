package domain

type User struct {
	Id           uint   `gorm:"primaryKey;column:Id"               json:"id"`
	Nombre       string `gorm:"size:100;not null;column:Nombre"    json:"nombre"`
	Username     string `gorm:"size:50;not null;unique;column:Username" json:"username"`
	Email        string `gorm:"size:120;not null;unique;column:Email"   json:"email"`
	PasswordHash string `gorm:"size:64;not null;column:PasswordHash"    json:"-"` // guardamos el hash, no la contraseña
	Rol          string `gorm:"type:enum('SOCIO','ADMIN');default:'SOCIO';column:Rol" json:"rol"`

	// Relaciones
	Inscripciones []Inscripcion `gorm:"foreignKey:IdUsuario" json:"inscripciones"`
}
