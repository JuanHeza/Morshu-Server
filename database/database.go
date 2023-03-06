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

type GeneralData struct {
	DataType  dt.DataType `json:"data_type,omitempty" bson:"data_type,omitempty"`
	Eliminado dt.Status   `json:"eliminado,omitempty" bson:"eliminado,omitempty"`
}

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
// Find all by criteria
func Find(criteria bson.M, collection string, output interface{}) (interface{}, error) {
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

//FindOne by criteria
func FindOne(criteria bson.M, collection string) (*mongo.SingleResult, error) {
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
	output := coll.FindOne(context.TODO(), criteria)
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
func DeleteOne(criteria bson.M, collection string) (interface{}, error) {
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
	result, err := coll.DeleteOne(context.TODO(), criteria)
	if err != nil {
		return nil, err
	}
	log.Printf("Documents Deleted: %v\n", result.DeletedCount)
	return result, nil
}
func DeleteMany(criteria bson.M, collection string, output interface{}) (interface{}, error) {
	return output, nil
}
func InsertOne(collection string, insert interface{}) (interface{}, error) {

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