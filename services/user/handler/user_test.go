package handler

import (
	"EvilPanda/services/user/model"
	dt "EvilPanda/util/dataType"
	"bytes"
	"encoding/json"
	"io/ioutil"
	_ "log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestLogIn(t *testing.T) {
	type args struct {
		method string
		path   string
	}
	tests := []struct {
		name     string
		args     args
		wantUser model.User
		wantCode int
	}{
		{name: "Bad Request", args: args{method: http.MethodGet, path: "/logInTest"}, wantUser: model.User{Username: "juan", Password: "heza"}, wantCode: http.StatusMethodNotAllowed},
		{name: "Empty Body", args: args{method: http.MethodPost, path: "/logInTest"}, wantCode: http.StatusBadRequest},
		{name: "Missing password", args: args{method: http.MethodPost, path: "/logInTest"}, wantUser: model.User{Username: "juan"}, wantCode: http.StatusBadRequest},
		{name: "Missing username", args: args{method: http.MethodPost, path: "/logInTest"}, wantUser: model.User{Password: "heza"}, wantCode: http.StatusBadRequest},
		{name: "user didnt exist", args: args{method: http.MethodPost, path: "/logInTest"}, wantUser: model.User{Username: "juancho", Password: "lagarto"}, wantCode: http.StatusUnauthorized},
		{name: "Good One", args: args{method: http.MethodPost, path: "/logInTest"}, wantUser: model.User{Username: "juan", Password: "heza"}, wantCode: http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := json.Marshal(tt.wantUser)
			if err != nil {
				t.Error("ERROR ON PARSING")
			}
			if tt.wantUser.Username == "" && tt.wantUser.Password == "" {
				user = []byte{}
			}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(tt.args.method, tt.args.path, bytes.NewReader(user))
			LogIn(w, r)
			if w.Result().StatusCode != tt.wantCode {
				res, _ := ioutil.ReadAll(w.Result().Body)
				t.Errorf("Status code returned, %d, did not match expected code %d || %v\n", w.Result().StatusCode, http.StatusOK, string(res))
			}

		})

	}
}

func TestLogOut(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "single",
			args: args{
				r: httptest.NewRequest(http.MethodGet, "/upper?word=abc", nil),
				w: httptest.NewRecorder(),
			},
		},
	}
	for _, tt := range tests {

		cookie := http.Cookie{
			Name:     os.Getenv("SESSION_KEY"),
			Value:    "someData",
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}
		tt.args.r.AddCookie(&cookie)
		t.Run(tt.name, func(t *testing.T) {
			LogOut(tt.args.w, tt.args.r)
		})
	}
}

func TestCrud(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name       string
		args       args
		wantedCode int
		body       model.User
	}{
		{
			name: "Get",
			args: args{
				r: httptest.NewRequest(http.MethodGet, "/user", nil),
				w: httptest.NewRecorder(),
			},
			wantedCode: http.StatusOK,
			body:       model.User{Id: "6405303725456574f5073a24"},
		},
		{
			name: "Delete",
			args: args{
				r: httptest.NewRequest(http.MethodDelete, "/user", nil),
				w: httptest.NewRecorder(),
			},
			wantedCode: http.StatusOK,
			body:       model.User{Id: "6405767ada1ef8e9312ff8d5"},
		},
		{
			name: "Put",
			args: args{
				r: httptest.NewRequest(http.MethodPut, "/user", nil),
				w: httptest.NewRecorder(),
			},
			wantedCode: http.StatusOK,
			body:       model.User{Id: "64057e43ce5570e7f8f51dd9", Username: "juan", Password: "heza", UserLevel: dt.Administrador_level},
		},
		// {
		// 	name: "Post",
		// 	args: args{
		// 		r: httptest.NewRequest(http.MethodPost, "/user", nil),
		// 		w: httptest.NewRecorder(),
		// 	},
		// 	wantedCode: http.StatusOK,
		// 	body:       model.User{Username: "arali", Password: "flores", UserLevel: dt.Administrador_level},
		// },
		{
			name: "Patch",
			args: args{
				r: httptest.NewRequest(http.MethodPatch, "/user", nil),
				w: httptest.NewRecorder(),
			},
			wantedCode: http.StatusMethodNotAllowed,
			body:       model.User{Username: "arali", Password: "flores", UserLevel: dt.Administrador_level},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := json.Marshal(tt.body)
			if err != nil {
				t.Error("ERROR ON PARSING")
			}
			tt.args.r.Body = ioutil.NopCloser(bytes.NewReader(user))
			Crud(tt.args.w, tt.args.r)
			if tt.args.w.Result().StatusCode != tt.wantedCode {
				res, _ := ioutil.ReadAll(tt.args.w.Result().Body)
				t.Errorf("Status code returned, %d, did not match expected code %d || %v\n", tt.args.w.Result().StatusCode, http.StatusOK, string(res))
			}
		})
	}
}
