package handler

import (
	// "EvilPanda/database"

	"EvilPanda/services/user/model"
	dt "EvilPanda/util/dataType"
	"encoding/json"
	"io/ioutil"

	// "EvilPanda/util"
	"errors"
	_ "log"
	"net/http"
)

/*
func LogIn(w http.ResponseWriter, r *http.Request) {
	err := database.SetSession(w, r)
	if Error(w, err) {
		http.Redirect(w, r, "/getProducts", http.StatusSeeOther)
	}
}
func LogOut(w http.ResponseWriter, r *http.Request){}

func Error(w http.ResponseWriter, err error) bool {
	if util.IsOk(err) {
		return true
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return false
} */
func LogIn(w http.ResponseWriter, r *http.Request) {
	var userModel model.User
	user, err := checkRequestError(w, r, http.MethodPost, userModel, []string{"Username", "Password"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	userModel = user.(model.User)
	userModel.Search()
	if userModel.UserLevel != dt.Invalid_level {

		data, err := json.Marshal(user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if len(data) != 0 {
			w.Write(data)
		}
	}
}

func LogOut(w http.ResponseWriter, r *http.Request) {}

func checkRequestError(w http.ResponseWriter, r *http.Request, methodAllowed string, dataType interface{}, fields []string) (data interface{}, err error) {
	var mapa map[string]interface{}
	if r.Method != methodAllowed {
		http.Error(w, "Bad Method", http.StatusMethodNotAllowed)
		err = errors.New("Bad Method")
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	if len(reqBody) == 0 {
		http.Error(w, "No Body", http.StatusBadRequest)
		err = errors.New("No Body")
		return
	}
	json.Unmarshal(reqBody, &mapa)
	for _, field := range fields {
		if val, ok := mapa[field]; !ok || val == "" {
			http.Error(w, "Missing Data", http.StatusBadRequest)
			err = errors.New("Missing Data")
			return
		}
	}
	json.Unmarshal(reqBody, &dataType)
	data = dataType
	return
}
