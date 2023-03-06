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
	Criteria_equal        = "$eq"
	Criteria_in           = "$in"
	Criteria_grater       = "$gt"
	Criteria_less         = "$lt"
	Criteria_grater_than  = "$gte"
	Criteria_less_than    = "$lte"
	Criteria_not_equal    = "$ne"
	Criteria_and          = "$and"
	Criteria_not          = "$not"
	Criteria_nor          = "$nor"
	Criteria_or           = "$or"
	Criteria_exists       = "$exists"
	Criteria_type         = "$type"
	Criteria_regex        = "$regex"
	Criteria_text         = "$text"
	Criteria_slice        = "$slice"

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
	Invalid_level UserLevel = iota
	Desarrollador_level
	Administrador_level
	Gerente_level
	Supervisor_level
	Trabajador_level
)

const (
	Invalid_Status Status = iota
	Eliminado
	No_Eliminado
)

var (
	Mongo_uri = os.Getenv("MONGODB_URI")
	Allow_Origin = os.Getenv("Allow_Origin")
	Database = os.Getenv("Database")
	CollectionNames = []string{Colleccion_cliente, Colleccion_producto, Colleccion_tienda, Colleccion_usuario, Colleccion_pedidos, Colleccion_inventario, Colleccion_ofertas}
)

type Criteria struct {
	Field       string
	Restriction string
	Value       interface{}
}
