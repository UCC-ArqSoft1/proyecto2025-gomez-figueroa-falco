package clients

import (
	"backend/dao"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	user := "root"
	password := ""
	host := "localhost"
	port := 3306
	database := "backend"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		user, password, host, port, database)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("error connecting to DB: %v", err))
	}

	DB.AutoMigrate(&dao.User{})
	DB.Create(&dao.User{
		ID:           1,
		Username:     "emiliano",
		PasswordHash: "121j212hs9812sj2189sj",
	})
}

func GetUserByUsername(username string) dao.User {
	var user dao.User
	// SELECT * FROM users WHERE username = ? LIMIT 1
	DB.First(&user, "username = ?", username)
	return user
}
