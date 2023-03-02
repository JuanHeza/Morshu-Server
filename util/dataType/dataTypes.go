package dataType

import "os"
type DataType int64
type UserLevel int64
const (
	Invalid_type DataType = iota
	Producto_type
    User_type
    Order_type
    Store_type
    ClientType

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

	Invalid_level UserLevel = iota
	Desarrollador_level
	Administrador_level
	Gerente_level
	Supervisor_level
	Trabajador_level

	Colleccion_cliente  = "clients"
	Colleccion_producto = "products"
	Colleccion_tienda   = "stores"
	Colleccion_usuario  = "users"
	Colleccion_pedidos  = "orders"
	Database_Name       = "EvilPanda"
)
var (
    Mongo_uri = os.Getenv("MONGODB_URI")
)

type Criteria struct {
	Field       string
	Restriction string
	Value       interface{}
}
