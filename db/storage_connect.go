package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sort"
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
		log.Fatal(err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	collections, err := client.Database("rest_api").ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	sort.Strings(collections)

	employeesCollectionIndex := sort.SearchStrings(collections, "employees")
	if collections[employeesCollectionIndex] != "employees" {
		if err := client.Database("rest_api").CreateCollection(context.TODO(), "employees"); err != nil {
			log.Fatal(err)
		}
		EmployeeCollection = client.Database("rest_api").Collection("employees")
	}else {
		EmployeeCollection = client.Database("rest_api").Collection("employees")
	}

	departamentsCollectionIndex := sort.SearchStrings(collections, "departaments")
	if collections[departamentsCollectionIndex] != "departaments" {
		if err := client.Database("rest_api").CreateCollection(context.TODO(), "departaments"); err != nil {
			log.Fatal(err)
		}
		DepartamentCollection = client.Database("rest_api").Collection("departaments")
	}else {
		DepartamentCollection = client.Database("rest_api").Collection("departaments")
	}
}
