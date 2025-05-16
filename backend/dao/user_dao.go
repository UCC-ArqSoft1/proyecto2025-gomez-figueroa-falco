package dao

type User struct {
	ID       int    `gorm:"primary_key"`
	Username string `gorm:"unique"`
	Password string `gorm:"notnull"`
}
