package controller

import (
	"database/sql"
	"fmt"

	"github.com/german0598/go-fiber/database"
	"github.com/german0598/go-fiber/model"
	"github.com/gofiber/fiber/v2"
)

func GetPedidos(c *fiber.Ctx) error {
	db := database.DB
	queryProductos := fmt.Sprintf(`select * from "Pedido_Detalle" pd, "Producto" p where p."Id" = pd."Producto_Id" and pd."Pedido_Id" = ?`)
	var (
		pedidos       []model.Pedido
		productos     []model.Producto
		rows          *sql.Rows
		errorConsulta error
		producto      model.Producto
	)
	db.Table("Pedido").Find(&pedidos)
	respuestaPedidos := make([]model.Pedido, len(pedidos))

	for _, pedido := range pedidos {
		rows, errorConsulta = db.Raw(queryProductos, pedido.Id).Rows()

		if errorConsulta != nil {
			fmt.Println("[error querySQL]", errorConsulta)
			continue
		}

		for rows.Next() {
			db.ScanRows(rows, &producto)
			productos = append(productos, producto)
		}
		// pedido.Productos = productos
		respuestaPedidos = append(respuestaPedidos, pedido)
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"data":    pedidos,
			"items":   len(pedidos),
			"status":  fiber.StatusOK,
			"success": true,
		},
	)
}

func GetPedido(c *fiber.Ctx) error {
	return c.SendString("Solo un Book")
}

func NewPedido(c *fiber.Ctx) error {
	c.Accepts("enctype/form-data")
	c.Accepts("multipart/form-data")
	c.Body()
	// file, _ := c.FormFile("document")
	// string := c.FormValue("Nombre")
	// // fmt.Println( c.Body() )
	// fmt.Println("[files]", file)
	// fmt.Println("[nombre]", string)
	// if err == nil {

	// }
	// fmt.Println("ALO POLICIA")

	return c.SendString("Nuevo Book")
}

func UpdatePedido(c *fiber.Ctx) error {
	return c.SendString("Actualizar Book")
}

func DeletePedido(c *fiber.Ctx) error {
	return c.SendString("Borrar Book")
}
