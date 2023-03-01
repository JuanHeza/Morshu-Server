package main

import (
	// pr "EvilPanda/services/product/handler"
	user "EvilPanda/services/user/handler"
	_ "bytes"
	// "encoding/gob"
	"fmt"
	_ "image"
	_ "image/jpeg"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	/* 	gob.Register(&User{})
	   	router.HandleFunc("/", web.LogIn).Methods("GET", "OPTIONS") */
	   	router.HandleFunc("/logIn", user.LogIn).Methods(http.MethodPost)
	//router.HandleFunc("/", ).Methods("")
	/* 	router.HandleFunc("/getProducts", pr.GetProducts).Methods("GET", "OPTIONS") */
	fmt.Println("Server Online")
	log.Fatal(http.ListenAndServe(":8080", router))
	// https://www.freecodecamp.org/news/how-to-build-a-web-app-with-go-gin-and-react-cffdc473576/
}

/*
func middlewareLog(handler http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
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
