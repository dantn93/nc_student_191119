package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang191119/nc_student/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

var Client *mongo.Client

const (
	DbName  = "golang11"
	ColName = "student"
)

func Test() interface{} {
	//
	fmt.Println("connect & insert db")
	return insertNumber()
}

func init() {

	connect()

}

func insertNumber() interface{} {
	collection := Client.Database("testing").Collection("numbers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		return nil
	}
	id := res.InsertedID
	return id
}

func connect() {
	log.Println("Try to connect mongo:", config.Config.Mongo.Uri)
	clientOptions := options.Client().ApplyURI(config.Config.Mongo.Uri)
	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetMinPoolSize(4)
	clientOptions.SetReadPreference(readpref.Nearest())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("ping error: %v", err)
	}
	Client = client
}
