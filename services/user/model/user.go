package model

import (
	db "EvilPanda/database"
	ct "EvilPanda/util/criteria"
	dt "EvilPanda/util/dataType"
	"fmt"
	_ "log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        string         `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string         `json:"username,omitempty" bson:"username,omitempty"`
	Password  string         `json:"password,omitempty" bson:"password,omitempty"`
	UserLevel dt.UserLevel   `json:"user_level,omitempty" bson:"user_level,omitempty"`
	Store     string         `json:"store,omitempty" bson:"store,omitempty"`
	General   db.GeneralData `bson:",inline"`
}

const (
	fields_Id        = "id"
	fields_Username  = "username"
	fields_Password  = "password"
	fields_UserLevel = "user_level"
	fields_Store     = "store"
)

//GetCheckList return a list of required values for given method
func GetCheckList(method string, single bool) ([]string, string) {
	switch method {
	case http.MethodDelete:
		return []string{fields_Id}, method
	case http.MethodGet:
		return []string{fields_Id}, method
		//return []string{fields_Id, fields_Store}, method
	case http.MethodPost:
		return []string{fields_Username, fields_UserLevel}, method
	case http.MethodPut:
		return []string{fields_Username, fields_UserLevel, fields_Id, fields_Password}, method
	default:
		return []string{}, ""
	}
}

func (us *User) Validate() {
	username := us.Username
	if username == "" {
		us.Username = "None"
	}
	password := us.Password
	if password == "" {
		us.Password = "None"
	}
	userlevel := us.UserLevel
	if userlevel == dt.Invalid_level {
		us.UserLevel = dt.Trabajador_level
	}
	store := us.Store
	if store == "" {
		us.Store = "None"
	}
	datatype := us.General.DataType
	if datatype == dt.Invalid_type {
		us.General.DataType = dt.User_type
	}
	eliminado := us.General.Eliminado
	if eliminado == dt.Invalid_Status {
		us.General.Eliminado = dt.No_Eliminado
	}
}

func (us *User) setCriteria(criteria []dt.Criteria) []dt.Criteria {
	if us.Username != "" {
		criteria = append(criteria, dt.Criteria{Field: "username", Value: us.Username})
	}
	if us.Password != "" {
		criteria = append(criteria, dt.Criteria{Field: "password", Value: us.Password})
	}
	if us.Store != "" {
		criteria = append(criteria, dt.Criteria{Field: "Store", Value: us.Store})
	}
	if us.Id != "" {
		id, err := primitive.ObjectIDFromHex(us.Id)
		if err == nil {
			criteria = append(criteria, dt.Criteria{Field: "_id", Value: id})
		}
	}
	return criteria
}

// Create a user
// Working
func (us *User) Create() bool {
	id, err := db.InsertOne(dt.Colleccion_usuario, us)
	if err != nil {
		return false
	}
	us.Id = fmt.Sprintf("%s", id)
	return true
}

// Read info of a single user
// Working
func (us *User) Read() bool {
	criteria := db.BuildCriteria(us.setCriteria(ct.Main_restriction([]dt.Criteria{}, dt.User_type)))
	results, err := db.FindOne(criteria, dt.Colleccion_usuario)
	if err != nil {
		return false
	}
	err = results.Decode(&us)
	return err != nil
}

func (us *User) ReadAll() (list []User) {
	return
}
func (us *User) Update() bool {
	return true
}
func (us *User) Deactivate() bool {
	return true
}

// Delete a single user, removes the data
// Working
func (us *User) Delete() bool {
	criteria := db.BuildCriteria(us.setCriteria(ct.Main_restriction([]dt.Criteria{}, dt.User_type)))
	_, err := db.DeleteOne(criteria, dt.Colleccion_usuario)
	return err != nil
}

// Search Check if the user already exist
// Working
func (us *User) Search() bool {
	criteria := db.BuildCriteria(us.setCriteria(ct.Main_restriction([]dt.Criteria{}, dt.User_type)))
	results, err := db.FindOne(criteria, dt.Colleccion_usuario)
	if err != nil {
		return false
	}
	err = results.Decode(&us)
	return err != nil 
}
