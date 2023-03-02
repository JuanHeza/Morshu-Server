package database

/*
import (
	db "EvilPanda/database"
	"EvilPanda/services/user/model"
	ct "EvilPanda/util/criteria"
	dt "EvilPanda/util/dataType"
	_ "log"
)

var (
	collection = "users"
)

func GetUser(us *model.User) *model.User {

	return us
}
func SetUser(us *model.User) *model.User {

	return us
}
func Update(us *model.User) *model.User {

	return us
}
func Create(us *model.User) *model.User {

	return us
}
func Read(us *model.User) *model.User {

	return us
}
func Deleter(us *model.User) *model.User {

	return us
}
func Search(us *model.User) (*model.User, error) {
	collection := dt.Colleccion_usuario
	criteria := db.BuildCriteria(ct.Main_restriction([]dt.Criteria{}, dt.User_type))
	results, err := db.FindOne(criteria, collection, us)
	if err != nil {
		return us, err
	}
	us = results.(*model.User)
	return us, nil
}
*/
