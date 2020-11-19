package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	EmployeeCollection    *mongo.Collection
	DepartamentCollection *mongo.Collection
)

// ConnectToMongo ...
func ConnectToMongo() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Println(err)
	}

	EmployeeCollection = client.Database("rest_api").Collection("employees")
	DepartamentCollection = client.Database("rest_api").Collection("departaments")

}
