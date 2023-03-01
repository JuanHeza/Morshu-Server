package handler

/*
import (
	db "EvilPanda/database"
	pr "EvilPanda/services/product/model"
	dt "EvilPanda/util/dataType"
	"encoding/json"
	"net/http"
	// "strconv"

	// "github.com/gorilla/mux"
)

// func GetProject(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	var payload = make(map[string]interface{})
// 	index, _ := strconv.Atoi(vars["id"])
// 	payload["Data"] = index

// 	json.NewEncoder(w).Encode(payload)
// }

func GetProducts(w http.ResponseWriter, r *http.Request) {
	criteria := []dt.Criteria{}
	results := pr.ReadCriteria(r, criteria)

	products, jsonError := json.Marshal(results)
	if db.Error(w, jsonError) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(products))
	}
}
*/
