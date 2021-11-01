package controller

import (
	"github.com/german0598/go-fiber/database"
	"github.com/german0598/go-fiber/model"
	"github.com/gofiber/fiber/v2"
)

func GetProductos(c *fiber.Ctx) error {
	db := database.DB
	var products []model.Producto
	db.Table("Producto").Find(&products)
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"items":   len(products),
			"status":  fiber.StatusOK,
			"success": true,
			"data":    products,
		},
	)
}

func GetProducto(c *fiber.Ctx) error {
	return c.SendString("Solo un Book")
}

func NewProducto(c *fiber.Ctx) error {
	return c.SendString("Nuevo Book")
}

func UpdateProducto(c *fiber.Ctx) error {
	return c.SendString("Actualizar Book")
}

func DeleteProducto(c *fiber.Ctx) error {
	return c.SendString("Borrar Book")
}
