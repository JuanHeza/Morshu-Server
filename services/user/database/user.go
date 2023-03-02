package database

import (
	dt "EvilPanda/util/dataType"
	ct "EvilPanda/util/criteria"
	"EvilPanda/services/user/model"
	_ "log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
)

var (
	collection = "users"
)

func GetUser(us *model.User) *model.User {

	return us
}
func SetUser(us *model.User) *model.User {

	return us
}
func Update(us *model.User) *model.User {

	return us
}
func Create(us *model.User) *model.User {

	return us
}
func Read(us *model.User) *model.User {

	return us
}
func Deleter(us *model.User) *model.User {

	return us
}
func Search(us *model.User) *model.User {
    	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dt.Mongo_uri))
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	database := dt.Database_Name
	collection := dt.Colleccion_usuario
	coll := client.Database(database).Collection(collection)
	criteria := ct.Main_restriction([]dt.Criteria{}, dt.User_type)
	buildedCriteria := db.BuildCriteria(criteria)
	cursor, err := coll.Find(context.TODO(), buildedCriteria)
	if err != nil {
		return
	}
	if err = cursor.All(context.TODO(), &list); err != nil {
		log.Fatal(err)
		return
	}
	database.Disconnect()
	return us
}
