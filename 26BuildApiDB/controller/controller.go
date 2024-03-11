package controller

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/princekumarg/monogoapi/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	godotenv.Load()
	connection := os.Getenv("MONGO_URI")
	clientOption := options.Client().ApplyURI(connection)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")

}
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}
