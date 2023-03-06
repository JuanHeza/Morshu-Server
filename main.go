package main

import (
	// pr "EvilPanda/services/product/handler"
	user "EvilPanda/services/user/handler"
	dt "EvilPanda/util/dataType"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	/* 	gob.Register(&User{})*/
	router.HandleFunc("/", user.Path).Methods(http.MethodGet)
	router.HandleFunc("/test", user.Test).Methods(http.MethodGet)
	router.HandleFunc("/logIn", user.LogIn).Methods(http.MethodPost)
	router.HandleFunc("/logOut", user.LogOut).Methods(http.MethodGet)
	router.HandleFunc("/user", user.Crud).Methods(http.MethodPost, http.MethodGet, http.MethodDelete, http.MethodPut)
	router.HandleFunc("/user/multiple", user.Crud).Methods(http.MethodPost, http.MethodGet, http.MethodDelete, http.MethodPut)

	//router.HandleFunc("/", ).Methods("")
	/* 	router.HandleFunc("/getProducts", pr.GetProducts).Methods("GET", "OPTIONS") */
	fmt.Println("Server Online")
	log.Fatal(http.ListenAndServe(":8080", router))
	// https://www.freecodecamp.org/news/how-to-build-a-web-app-with-go-gin-and-react-cffdc473576/
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", dt.Allow_Origin)
}
/*
func middlewareLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			enableCors(&w)
			loggedIn := db.LoggedIn(r)
			if loggedIn.Username != "" {
				handler.ServeHTTP(w, r)
			} else {
				http.Redirect(w, r, "/logIn", http.StatusSeeOther)
			}
			log.Printf("Nueva petición. Método: %s. IP: %s. URL solicitada: %s\n", r.Method, r.RemoteAddr, r.URL)
		})
}
*/
