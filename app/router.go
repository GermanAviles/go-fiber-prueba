package app

import (
	"github.com/german0598/go-fiber/controller"
	"github.com/gofiber/fiber/v2"
)

func configurarRutas(app *fiber.App) {
	group := agruparRutas(app)
	books(group)
	shoes(group)
	pedido(group)
	producto(group)
}

func books(group fiber.Router) {

	group.Post("/book", controller.NewBook)
	group.Get("/books", controller.GetBooks)
	group.Get("/book/:id", controller.GetBook)
	group.Patch("/book/:id", controller.UpdateBook)
	group.Delete("/book/:id", controller.DeleteBook)
}

func shoes(group fiber.Router) {

	group.Post("/shoe", controller.NewShoe)
	group.Get("/shoes", controller.GetShoes)
	group.Get("/shoe/:id", controller.GetShoe)
	group.Patch("/shoe/:id", controller.UpdateShoe)
	group.Delete("/shoe/:id", controller.DeleteShoe)
}

func pedido(group fiber.Router) {

	group.Post("/pedido", controller.NewPedido)
	group.Get("/pedidos", controller.GetPedidos)
	group.Get("/pedido/:id", controller.GetPedido)
	group.Patch("/pedido/:id", controller.UpdatePedido)
	group.Delete("/pedido/:id", controller.DeletePedido)
}

func producto(group fiber.Router) {

	group.Post("/producto", controller.NewProducto)
	group.Get("/productos", controller.GetProductos)
	group.Get("/producto/:id", controller.GetProducto)
	group.Patch("/producto/:id", controller.UpdateProducto)
	group.Delete("/producto/:id", controller.DeleteProducto)
}

func agruparRutas(app *fiber.App) fiber.Router {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	return v1
}
