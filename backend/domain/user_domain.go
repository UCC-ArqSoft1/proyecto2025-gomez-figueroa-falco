package domain

type User struct {
	Id           uint   `json:"id"`
	Nombre       string `json:"nombre"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"` // guardamos el hash, no la contraseña
	Rol          string `json:"rol"`

	// Relaciones
	Inscripciones []Inscripcion `json:"inscripciones"`
}
