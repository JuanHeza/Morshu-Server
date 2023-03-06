package dataType

import "os"

type DataType int8
type UserLevel int8
type Status int8

const (
	Invalid_type DataType = iota
	Producto_type
	User_type
	Order_type
	Store_type
	ClientType
)

const (
	Criteria_equal       string = "$eq"
	Criteria_in          string = "$in"
	Criteria_grater      string = "$gt"
	Criteria_less        string = "$lt"
	Criteria_grater_than string = "$gte"
	Criteria_less_than   string = "$lte"
	Criteria_not_equal   string = "$ne"
	Criteria_and         string = "$and"
	Criteria_not         string = "$not"
	Criteria_nor         string = "$nor"
	Criteria_or          string = "$or"
	Criteria_exists      string = "$exists"
	Criteria_type        string = "$type"
	Criteria_regex       string = "$regex"
	Criteria_text        string = "$text"
	Criteria_slice       string = "$slice"
)

const (
	Invalid_level UserLevel = iota
	Desarrollador_level
	Administrador_level
	Gerente_level
	Supervisor_level
	Trabajador_level
)

const (
	Colleccion_cliente    = "clientes"
	Colleccion_producto   = "productos"
	Colleccion_tienda     = "tienda"
	Colleccion_usuario    = "usuarios"
	Colleccion_pedidos    = "pedidos"
	Colleccion_inventario = "inventario"
	Colleccion_ofertas    = "ofertas"
	Database_Name         = "pruebas"
)

const (
	Invalid_Status Status = iota
	Eliminado
	No_Eliminado
)

var (
	Mongo_uri = os.Getenv("MONGODB_URI")
)

type Criteria struct {
	Field       string
	Restriction string
	Value       interface{}
}
