package main
import (
	_"bytes"
	"fmt"
	_"image"
	_"image/jpeg"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
		"github.com/gorilla/sessions"
    "context"
	"encoding/json"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var store  = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", setCookie).Methods("GET", "OPTIONS")
	router.HandleFunc("/getProducts", GetProducts).Methods("GET", "OPTIONS")
	router.HandleFunc("/Reporte/", GetProject).Methods("GET", "OPTIONS")
	router.HandleFunc("/Correo/", GetProject).Methods("GET", "OPTIONS")
	return router
}

func main() {
    //insertItem(os.Getenv("MONGODB_URI"))
	r := Router()
	fmt.Println("Server Online")
	log.Fatal(http.ListenAndServe(":8080", r))
	// https://www.freecodecamp.org/news/how-to-build-a-web-app-with-go-gin-and-react-cffdc473576/
}


func GetProject(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var payload = make(map[string]interface{})
	index, _ := strconv.Atoi(vars["id"])
	payload["Data"] = index
	
	json.NewEncoder(w).Encode(payload)
}
    
func MongoConection(uri string, r *http.Request)(response []byte){
    session, _ := store.Get(r, "session-name")
    database := fmt.Sprintf("%v",session.Values["database"])
    collection := fmt.Sprintf("%v",session.Values["collection"])
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	//uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(database).Collection(collection)
	var results []Product
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
    /*
	for _, result := range results {
		response += result.Print()
        //fmt.Printf("%+v\n", result)
	}
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", "title")
		return
	}
 */
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

func insertItem(uri string)(int){
    item := New("Pulsera Bolita", "Pulserra", 10, false, 100, 12, nil, nil)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("pruebas").Collection("tienda")
    fmt.Println(*item)
    
    result, err := coll.InsertOne(context.TODO(), item)
    if err != nil {
    	panic(err)
        return 1
    }
    fmt.Println(result)
    return 0
}
func setCookie(w http.ResponseWriter, r *http.Request){
    session, _ := store.Get(r, "session-name")
    session.Values["database"] = "pruebas"
    session.Values["collection"] = "tienda"
    err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
    
}
func GetProducts(w http.ResponseWriter, r *http.Request){
    session, _ := store.Get(r, "session-name")
    fmt.Println(session.Values["collection"])
    products := MongoConection(os.Getenv("MONGODB_URI"), r)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(products))
}