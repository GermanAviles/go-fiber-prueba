package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/german0598/go-fiber/app"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/german0598/go-fiber/config"
	"github.com/german0598/go-fiber/database"
)

func main() {
	// Cargamos las variables de entorno
	cargarDotEnv()
	// Instanciamos (creamos) una variable de tipo servidor
	servidor := app.Servidor{
		App:    fiber.New(),
		Puerto: config.GetEnv("SERVER_PORT"),
	}
	// Middleware para ver por consola las peticiones al servidor
	servidor.App.Use(logger.New())
	servidor.App.Use(cors.New())
	// Abrimos conexión con la base de datos
	database.ConnectDB()
	// Cerramos conexión con la base de datos
	defer closeConnectionDB()
	// Configuración de los endpoints de mi API
	servidor.IniciarRutas()
	// Inicialización del servidor
	servidor.IniciarServidor()
}

func closeConnectionDB() {
	db, err := database.DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}

func cargarDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
