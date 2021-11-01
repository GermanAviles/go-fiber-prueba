package database

import (
	"fmt"
	"strconv"

	"github.com/german0598/go-fiber/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB gorm connector
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	puerto := config.GetEnv("DB_PORT")
	port, err := strconv.ParseUint(puerto, 10, 32)
	fmt.Println("[PUERTO DB] ", port)
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.GetEnv("DB_HOST"),
		port,
		config.GetEnv("DB_USERNAME"),
		config.GetEnv("DB_PASSWORD"),
		config.GetEnv("DB_NAME"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("[ERROR] Fallo al conectar con la base de datos...")
	}

	fmt.Println("[Postgresql] Conexión establecida exitosamente")
}
