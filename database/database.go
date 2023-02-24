package database

import (
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

type MongoClient struct {
	user, password string
	collection     []string
}

var (
	user   = MongoClient{}
	client *mongo.Client
	Store  = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
)

func getSession(r *http.Request, search string) string {
	session, _ := Store.Get(r, "session-name")

	return fmt.Sprintf("%v", session.Values[search])
}
func SetSession(w http.ResponseWriter, r *http.Request) (err error) {
	session, _ := Store.Get(r, "session-name")
	session.Values["database"] = os.Getenv("TEST_DATABASE")
	session.Values["collection"] = os.Getenv("TEST_COLLECTION")
	session.Values["user"] = os.Getenv("TEST_USER")
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
	return
}

func GetCollection(r *http.Request) (coll *mongo.Collection) {
	if client == nil {
		client = user.connect(os.Getenv("MONGODB_URI"))
	}
	/*defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()*/
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

/*
import "go.mongodb.org/mongo-driver/mongo"
serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
clientOptions := options.Client().
    ApplyURI("mongodb+srv://JuanHeza:1hCYw6lH9fF26Prs@evilpanda.cgorqpw.mongodb.net/?retryWrites=true&w=majority").
    SetServerAPIOptions(serverAPIOptions)
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, clientOptions)
if err != nil {
    log.Fatal(err)
}
*/
