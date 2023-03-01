package dataType

type DataType int64
type UserLevel int64

const (
	_ DataType = iota
	Producto_type

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

	_ UserLevel = iota
	Desarrollador_level
	Administrador_level
	Gerente_level
	Supervisor_level
	Trabajador_level
)

type Criteria struct {
	Field       string
	Restriction string
	Value       interface{}
}

