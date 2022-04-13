package models

import (
	"context"
	"log"
	"os"

	"github.com/DumDumGeniuss/game-of-liberty-computer/models/gameblockmodel"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var err error

func InitializeModels() {
	uri := os.Getenv("MONGODB_URI")
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	database := os.Getenv("MONGODB_DATABASE")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	if username == "" {
		log.Fatal("You must set your 'MONGODB_USERNAME'")
	}
	if password == "" {
		log.Fatal("You must set your 'MONGODB_PASSWORD'")
	}
	if database == "" {
		log.Fatal("You must set your 'MONGODB_DATABASE'")
	}
	// credential := options.Credential{
	// 	Username: username,
	// 	Password: password,
	// }
	client, err = mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(uri), //.SetAuth(credential),
	)
	if err != nil {
		panic(err)
	}
	// if err := client.Disconnect(context.TODO()); err != nil {
	// 	panic(err)
	// }
	gameblockmodel.IniitalizeGameBlocks(client.Database(database).Collection("GameBlock"))
}
