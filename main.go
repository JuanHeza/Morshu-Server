package main

import (
	web "EvilPanda/controller"
	_ "bytes"
	"fmt"
	_ "image"
	_ "image/jpeg"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", web.LogIn).Methods("GET", "OPTIONS")
	router.HandleFunc("/getProducts", web.GetProducts).Methods("GET", "OPTIONS")
	fmt.Println("Server Online")
	log.Fatal(http.ListenAndServe(":8080", router))
	// https://www.freecodecamp.org/news/how-to-build-a-web-app-with-go-gin-and-react-cffdc473576/
}

/*
func MongoConection(uri string, r *http.Request) (response []byte) {
	var results []Product
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	//
			for _, result := range results {
				response += result.Print()
		        //fmt.Printf("%+v\n", result)
			}
			if err == mongo.ErrNoDocuments {
				fmt.Printf("No document was found with the title %s\n", "title")
				return
			}
	//
	response, jsonError := json.Marshal(results)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}
	if err != nil {
		panic(err)
	}
	//jsonData, err := json.MarshalIndent(results, "", "    ")
	//	if err != nil {
	//		panic(err)
	//}
	//fmt.Printf("%s\n", jsonData)
	fmt.Println("END")
	return
}
*/
