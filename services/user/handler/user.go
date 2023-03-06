package handler

import (
	// "EvilPanda/database"

	"EvilPanda/services/user/model"
	dt "EvilPanda/util/dataType"
	ss "EvilPanda/util/session"
	"encoding/json"
	"io/ioutil"
	"log"

	// "EvilPanda/util"
	"errors"
	_ "fmt"
	"net/http"
)

func LogIn(w http.ResponseWriter, r *http.Request) {
	var userModel model.User
	user, err := checkRequestError(w, r, http.MethodPost, userModel, []string{"username", "password"})
	if err != nil {
		return
	}
	json.Unmarshal(user, &userModel)
	found := userModel.Search()
	if found && userModel.UserLevel != dt.Invalid_level {
		ss.SetCookieHandler(&w, r, userModel)
		data, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if len(data) != 0 {
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		}
	} else {
		http.Error(w, "Unauthorized ", http.StatusUnauthorized)
	}
}

func Test(w http.ResponseWriter, r *http.Request) {
	var userModel model.User = model.User{Username: "juan", Password: "heza"}
	found := userModel.Search()
	if found && userModel.UserLevel != dt.Invalid_level {
		ss.SetCookieHandler(&w, r, userModel)
		cookie := http.Cookie{
			Name:     "example-Cookie",
			Value:    "Hello world!",
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(w, &cookie)
		log.Println(w.Header())
		data, err := json.Marshal(userModel)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if len(data) != 0 {
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		}
	}
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	ss.CleanCookieHandler(&w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Path(w http.ResponseWriter, r *http.Request) {
	ss.GetCookieHandler(w, r)
	log.Println(r.Cookies())
}

func Crud(w http.ResponseWriter, r *http.Request) {
	var userModel model.User
	checkList, method := model.GetCheckList(r.Method, true)
	user, err := checkRequestError(w, r, method, userModel, checkList)
	if err != nil {
		return
	}
	json.Unmarshal(user, &userModel)
	//userModel.Validate()
	switch r.Method {
	case http.MethodPost:
		res := userModel.Create()
		if !res {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodGet:
		res := userModel.Read()
		if !res {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPut:
		res := userModel.Update()
		if !res {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodDelete:
		res := userModel.Delete()
		if !res {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(userModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(data)
}

func checkRequestError(w http.ResponseWriter, r *http.Request, methodAllowed string, dataType interface{}, fields []string) (data []byte, err error) {
	var mapa map[string]interface{}
	if r.Method != methodAllowed {
		http.Error(w, "Bad Method", http.StatusMethodNotAllowed)
		err = errors.New("error: Bad Method")
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	if len(reqBody) == 0 {
		http.Error(w, "No Body", http.StatusBadRequest)
		err = errors.New("error: No Body")
		return
	}
	json.Unmarshal(reqBody, &mapa)
	for _, field := range fields {
		if val, ok := mapa[field]; !ok || val == "" {
			http.Error(w, "Missing Data", http.StatusBadRequest)
			err = errors.New("error: Missing Data")
			return
		}
	}
	data = reqBody
	return
}
