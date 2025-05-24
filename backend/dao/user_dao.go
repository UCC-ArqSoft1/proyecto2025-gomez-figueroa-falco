package dao

type User struct {
	Id           uint   `gorm:"primaryKey:autoIncrement;column:id"`
	Nombre       string `gorm:"size:100;not null;column:nombre"`
	Username     string `gorm:"size:50;not null;unique;column:username"`
	Email        string `gorm:"size:120;not null;unique;column:email"`
	PasswordHash string `gorm:"size:64;not null;column:password_hash"` // guardamos el hash, no la contraseña
	Rol          string `gorm:"type:enum('SOCIO','ADMIN');default:'SOCIO';column:rol"`
}
