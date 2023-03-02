package criteria

import dt "EvilPanda/util/dataType"

func Main_restriction(criteria []dt.Criteria, datatype dt.DataType) []dt.Criteria {
	criteria = append(criteria, dt.Criteria{Field: "data_type", Value: datatype})
	criteria = append(criteria, dt.Criteria{Field: "eliminado", Value: false})
	return criteria
}