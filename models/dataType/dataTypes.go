package dataType

type DataType int64

const (
	Producto_type DataType = iota

	Criteria_equal       string = "equals"
	Criteria_contains    string = "contains"
	Criteria_grater      string = "greater"
	Criteria_less        string = "less"
	Criteria_grater_than string = "greater than"
	Criteria_less_than   string = "less than"
	Criteria_includes    string = "includes"
	Criteria_not         string = "not"
)

type Criteria struct {
	Field       string
	Restriction string
	Value       interface{}
}
