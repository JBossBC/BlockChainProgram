package dao

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

//这里需要优化
func initDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://mongo:6379")
	var err error
	MongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(fmt.Errorf("mongo db init err :%s", err.Error()))
	}
}
