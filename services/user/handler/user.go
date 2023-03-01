package handler

import (
	// "EvilPanda/database"

	"EvilPanda/services/user/model"
	"encoding/json"
	"io/ioutil"

	// "EvilPanda/util"
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
	var user model.User
	data := checkRequestError(w, r, http.MethodPost, user, []string{"Username", "Password"})
	if len(data) != 0 {
		w.Write(data)
	}
}

func LogOut(w http.ResponseWriter, r *http.Request) {}

func checkRequestError(w http.ResponseWriter, r *http.Request, methodAllowed string, dataType interface{}, fields []string) (data []byte) {
	var mapa map[string]interface{}
	if r.Method != methodAllowed {
		http.Error(w, "Bad Method", http.StatusMethodNotAllowed)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	if len(reqBody) == 0 {
		http.Error(w, "No Body", http.StatusBadRequest)
		return
	}
	json.Unmarshal(reqBody, &mapa)
	for _, field := range fields {
		if val, ok := mapa[field]; !ok || val == "" {
			http.Error(w, "Missing Data", http.StatusBadRequest)
			return
		}
	}
	json.Unmarshal(reqBody, &dataType)
	data, err := json.Marshal(dataType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}
