package model

import (
	"time"
)

type ProductosPedido struct {
	Id              int       `gorm:"primaryKey" json:"id"`
	Producto_Id     uint      `gorm:"not null" json:"productoId"`
	Pedido_Id       uint      `gorm:"not null" json:"pedidoId"`
	Nombre_Producto string    `gorm:"not null" json:"nombreProducto"`
	Cantidad        float32   `gorm:"not null" json:"cantidad"`
	Vlr_Unit        float32   `gorm:"not null" json:"valorUnit"`
	Fecha_Creacion  time.Time `json:"fechaCreacion"`
}
