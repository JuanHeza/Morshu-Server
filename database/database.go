package database

import (
	dt "EvilPanda/util/dataType"
	us "EvilPanda/services/user/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
//	user   = MongoClient{}
//	client *mongo.Client
//	Store  = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
)
/*

type MongoClient struct {
	user, password string
	collection     []string
}


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

func (mc *MongoClient) connect(uri string) (client *mongo.Client) {
	// user, password, cluster string
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return
}
func Disconnect() (err error) {
	err = client.Disconnect(context.TODO())
	client = nil
	return
}

func GetCollection(r *http.Request) (coll *mongo.Collection) {
	if client == nil {
		client = user.connect(os.Getenv("MONGODB_URI"))
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	database := getSession(r, "database")
	collection := getSession(r, "collection")
	coll = client.Database(database).Collection(collection)
	return
}

func MongoConection(uri string) {
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
