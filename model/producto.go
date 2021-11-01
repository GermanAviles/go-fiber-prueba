package model

import (
	"time"
)

type Producto struct {
	Id                    int       `gorm:"primaryKey" json:"id"`
	Proveedor_id          string    `gorm:"not null" json:"provedorId"`
	Nombre                string    `gorm:"not null" json:"nombre"`
	Presentacion          string    `gorm:"not null" json:"presentacion"`
	Descripcion           string    `gorm:"not null" json:"descripcion"`
	Cantidad              string    `gorm:"not null" json:"cantidad"`
	Vlr_Unit              string    `gorm:"not null" json:"valorUnit"`
	Prc_Impuesto          string    `gorm:"not null" json:"precioImpuesto"`
	Prc_Descuento         string    `gorm:"not null" json:"precioDescuento"`
	Es_Destacado          bool      `gorm:"not null" json:"esDestacado"`
	Tipo_Unidad           uint      `gorm:"not null" json:"tipoUnidad"`
	Peso_Kg               float32   `gorm:"not null" json:"pesoKg"`
	Activo                bool      `gorm:"not null" json:"activo"`
	Fecha_Creacion        time.Time `gorm:"not null" json:"fechaCreacion"`
	Categoria_Producto_Id uint      `json:"categoriaProductoId"`
	Ruta_Ficha_Tecnica    string    `json:"rutaFichaTecnica"`
	Fecha_Modificacion    time.Time `json:"fechaModificacion"`
}
