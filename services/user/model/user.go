package model

import "EvilPanda/util/dataType"

type User struct{
	Username string
	Password string
	SessionExpire string
	UserLevel dataType.UserLevel
}

func (us *User) SetSession()(){}
func (us *User) GetSession()(){}
func (us *User) Create()(){}
func (us *User) Read()(){}
func (us *User) Update()(){}
func (us *User) Delete()(){}
func (us *User) Search()(found bool){
    return
}