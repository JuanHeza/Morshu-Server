package database

import (
	dt "EvilPanda/util/dataType"
	"context"
	_ "encoding/json"
	_ "fmt"
	"log"
	_ "net/http"
	_ "os"

	_ "github.com/gorilla/sessions"

	_ "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
func LoggedIn(r *http.Request) (actual us.User) {
	session, _ := Store.Get(r, "session-name")
	actual.UserLevel = session.Values["userlevel"].(dt.UserLevel)
	actual.Username = fmt.Sprintf("%v", session.Values["username"])
	return
}
func DestroySession(r *http.Request, w http.ResponseWriter) (err error) {
	session, _ := Store.Get(r, "session-name")
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	return
}
func getSession(r *http.Request, search string) string {
	session, _ := Store.Get(r, "session-name")

	return fmt.Sprintf("%v", session.Values[search])
}
func SetSession(w http.ResponseWriter, r *http.Request) (err error) {
	session, _ := Store.Get(r, "session-name")
	session.Values["database"] = os.Getenv("TEST_DATABASE")
	session.Values["collection"] = os.Getenv("TEST_COLLECTION")
	session.Values["username"] = os.Getenv("TEST_USER")
	session.Values["userlevel"] = dt.Desarrollador_level
	err = session.Save(r, w)
	return
}

func MongoConection(uri string) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("sample_mflix").Collection("movies")
	title := "Back to the Future"
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
*/
func BuildCriteria(criteria []dt.Criteria) (multi bson.M) {
	//var one interface{}
	multi = bson.M{}
	for _, ctr := range criteria {
		if ctr.Restriction == "" {
			multi[ctr.Field] = ctr.Value
		} else {
			multi[ctr.Field] = bson.D{{ctr.Restriction, ctr.Value}}
		}
	}
	return
}
func Find(criteria bson.M, collection string, output interface{}) (interface{}, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dt.Mongo_uri))
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(dt.Database_Name).Collection(collection)
	cursor, err := coll.Find(context.TODO(), criteria)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &output); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return output, nil
}

func FindOne(criteria bson.M, collection string, output interface{}) (interface{}, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dt.Mongo_uri))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(dt.Database_Name).Collection(collection)
	err = coll.FindOne(context.TODO(), criteria).Decode(&output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func FindOneAndReplace(criteria bson.M, collection string, output interface{}) (interface{}, error) {
	return output, nil
}
func FindOneAndUpdate(criteria bson.M, collection string, output interface{}) (interface{}, error) {
	return output, nil
}
func FindOneAndDelete(criteria bson.M, collection string, output interface{}) (interface{}, error) {
	return output, nil
}
func DeleteOne(criteria bson.M, collection string, output interface{}) (interface{}, error) {
	return output, nil
}
func DeleteMany(criteria bson.M, collection string, output interface{}) (interface{}, error) {
	return output, nil
}
func InsertOne(criteria bson.M, collection string, insert interface{}) (interface{}, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dt.Mongo_uri))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// begin insertOne
	coll := client.Database(dt.Database_Name).Collection(collection)
	result, err := coll.InsertOne(context.TODO(), insert)
	if err != nil {
		//panic(err)
		return nil, err
	}

	log.Printf("Document inserted with ID: %s\n", result.InsertedID)
	return result, nil
}
func InsertMany(criteria bson.M, collection string, output interface{}) (interface{}, error) {
	return output, nil
}
func ReplaceOne(criteria bson.M, collection string, output interface{}) (interface{}, error) {
	return output, nil
}
func UpdateById(criteria bson.M, collection string, output interface{}) (interface{}, error) {
	return output, nil
}
func UpdateOne(criteria bson.M, collection string, output interface{}) (interface{}, error) {
	return output, nil
}
func UpdateMany(criteria bson.M, collection string, output interface{}) (interface{}, error) {
	return output, nil
}
func Aggregate(criteria bson.M, collection string, output interface{}) (interface{}, error) {
	return output, nil
}
func Distinct(criteria bson.M, collection string, output interface{}) (interface{}, error) {
	return output, nil
}

//
// import "go.mongodb.org/mongo-driver/mongo"
// serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
// clientOptions := options.Client().
//     ApplyURI("MONGO_STRING").
//     SetServerAPIOptions(serverAPIOptions)
// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// defer cancel()
// client, err := mongo.Connect(ctx, clientOptions)
// if err != nil {
//     log.Fatal(err)
// }
//
