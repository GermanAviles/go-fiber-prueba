package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("Te muestro todos los Books")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("Solo un Book")
}

func NewBook(c *fiber.Ctx) error {
	return c.SendString("Nuevo Book")
}

func UpdateBook(c *fiber.Ctx) error {
	return c.SendString("Actualizar Book")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Borrar Book")
}
