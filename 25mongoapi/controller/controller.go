package controller

import (
	"context"
	"fmt"
	"log"

	model "command-line-argumentsC:\\Users\\rajde\\OneDrive\\Desktop\\Go Learn\\mygolang-lco\\25mongoapi\\model\\models.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const connectionString = "mongodb+srv://rajdeepmallick010:rajdeep@golearn.ltqa1rp.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"
const colName = "watchlist"
 
// Most important
var collection *mongo.Collection

// connect with mongodb
func init() {
	// provide the client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
}

// MongoDB Helpers


// insert 1 record

//  we have define the package as model so we wrote "model"
func insertOneMovie(movie model.Netflix) {

	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}


// update 1 record

func updateOneMovie(movieId string){
	primitive.ObjectIDFromHex(movieId)
}