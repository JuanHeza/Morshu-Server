package model

import (
	db "EvilPanda/database"
	ct "EvilPanda/util/criteria"
	dt "EvilPanda/util/dataType"
	"log"
)

type User struct {
	Username      string
	Password      string
	SessionExpire string
	UserLevel     dt.UserLevel
}

func (us *User) SetSession() {}
func (us *User) GetSession() {}
func (us *User) Create()     {}
func (us *User) Read()       {}
func (us *User) Update()     {}
func (us *User) Delete()     {}
func (us *User) Search() bool {
	collection := dt.Colleccion_usuario
	criteria := db.BuildCriteria(ct.Main_restriction([]dt.Criteria{}, dt.User_type))
	results, err := db.FindOne(criteria, collection, us)
	if err != nil {
		return false
	}
	us = results.(*User)
	log.Println(us)
	return true
}
