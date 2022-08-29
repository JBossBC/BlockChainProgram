package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

//这里需要优化
func InitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	clientOptions.SetAuth(options.Credential{Username: "root", Password: "root"})
	var err error
	MongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(fmt.Errorf("mongo db init err :%s", err.Error()))
	}
	log.Println("Init mongoDB successfully")
}
