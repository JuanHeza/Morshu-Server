package controller

import (
	"EvilPanda/database"
	"EvilPanda/util"
	"net/http"
)

func LogIn(w http.ResponseWriter, r *http.Request) {
	err := database.SetSession(w, r)
	if Error(w, err) {
		http.Redirect(w, r, "/getProducts", http.StatusSeeOther)
	}
}

func Error(w http.ResponseWriter, err error) bool {
	if util.IsOk(err) {
		return true
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return false
}
