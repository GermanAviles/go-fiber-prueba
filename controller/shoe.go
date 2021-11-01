package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetShoes(c *fiber.Ctx) error {
	return c.SendString("Te muestro todos los Zapatos")
}

func GetShoe(c *fiber.Ctx) error {
	return c.SendString("Solo un Zapato")
}

func NewShoe(c *fiber.Ctx) error {
	return c.SendString("Nuevo Zapato")
}

func UpdateShoe(c *fiber.Ctx) error {
	return c.SendString("Actualizar Zapato")
}

func DeleteShoe(c *fiber.Ctx) error {
	return c.SendString("Borrar Zapato")
}
