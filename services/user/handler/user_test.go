package handler

import (
	"EvilPanda/services/user/model"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
	}{
		{name: "Bad Request", args: args{method: http.MethodGet, path: "/logInTest"}, wantUser: model.User{Username: "juan", Password: "heza"}},
		{name: "Empty Body", args: args{method: http.MethodPost, path: "/logInTest"}},
		{name: "Missing password", args: args{method: http.MethodPost, path: "/logInTest"}, wantUser: model.User{Username: "juan"}},
		{name: "Missing username", args: args{method: http.MethodPost, path: "/logInTest"}, wantUser: model.User{Password: "heza"}},
		{name: "user didnt exist", args: args{method: http.MethodPost, path: "/logInTest"}, wantUser: model.User{Username: "juancho", Password: "lagarto"}},
		{name: "Good One", args: args{method: http.MethodPost, path: "/logInTest"}, wantUser: model.User{Username: "juan", Password: "heza"}},
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
			if w.Result().StatusCode != http.StatusOK {
				res, _ := ioutil.ReadAll(w.Result().Body)
				t.Errorf("Status code returned, %d, did not match expected code %d || %v\n", w.Result().StatusCode, http.StatusOK, string(res))
			}
		})

	}
}
