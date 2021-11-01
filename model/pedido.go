package model

type Pedido struct {
	Id                  int         `gorm:"primaryKey"`
	Productos           []*Producto `json:"productos,omitempty"`
	Usuario_Id          uint        `gorm:"not null" json:"usuarioId"`
	Estado_Pedido       uint        `json:"estadoPedido"`
	Fecha_Estado_Pedido string      `json:"fechaEstadoPedido"`
	Estado_Pago_Id      int         `json:"estadoPagoId"`
	Fecha_Estado_Pago   string      `json:"fechaEstadoPago"`
	Fecha_Estimada      string      `json:"fechaEstimada"`
	Ubicacion_Id        uint        `gorm:"not null" json:"ubicacionId"`
	Direccion_Entrega   string      `json:"direccionEntrega"`
	Vlr_Envio           float32     `gorm:"not null" json:"valorEnvio"`
	Vlr_Descuento       float32     `gorm:"not null" json:"valorDescuento"`
	Vlr_Impuesto        float32     `gorm:"not null" json:"valorImpuesto"`
	Vlr_Total           float32     `gorm:"not null" json:"valorTotal"`
	Contacto_Entrega    string      `gorm:"not null" json:"contactoEntrega"`
	Telefono_Contacto   string      `gorm:"not null" json:"telefonoContacto"`
	Tipo_Pago           uint        `gorm:"not null" json:"tipoPago"`
	Ruta_Pago           string      `json:"rutaPago"`
	Info_Pago           string      `json:"infoPago"`
	Observaciones       string      `json:"observaciones"`
	Fecha_Creacion      string      `gorm:"not null" json:"fechaCreacion"`
}
