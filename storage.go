package main

import (
	"context"
	"errors"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Employee ...
type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

// Storage ...
type Storage interface {
	Insert(e *Employee) error
	Get(id int) (Employee, error)
	Update(id int, e *Employee) error
	Delete(id int) error
	GetAll() (map[int]Employee, error)
}

// MongoStorage ...
type MongoStorage struct {
	counter int
	sync.Mutex
}

// NewMongoStorage ...
func NewMongoStorage() *MongoStorage {
	return &MongoStorage{
		counter: 1,
	}
}

// Insert ...
func (s *MongoStorage) Insert(e *Employee) error {
	s.Lock()
	defer s.Unlock()

	e.ID = s.counter
	_, err := Collection.InsertOne(context.TODO(), bson.M{
		"_id":    *&e.ID,
		"name":   *&e.Name,
		"sex":    *&e.Sex,
		"age":    *&e.Age,
		"salary": *&e.Salary,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	s.counter++

	return nil
}

// Get ...
func (s *MongoStorage) Get(id int) (Employee, error) {
	s.Lock()
	defer s.Unlock()

	var employee Employee

	employee.ID = id

	err := Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&employee)
	if err != nil {
		log.Println(err)
		return employee, err
	}

	return employee, nil
}

// Update ...
func (s *MongoStorage) Update(id int, e *Employee) error {
	s.Lock()
	defer s.Unlock()

	_, err := Collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{{"$set",
			bson.M{
				"_id":    id,
				"name":   *&e.Name,
				"sex":    *&e.Sex,
				"age":    *&e.Age,
				"salary": *&e.Salary,
			}}})
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

// Delete ...
func (s *MongoStorage) Delete(id int) error {
	s.Lock()
	defer s.Unlock()

	deleteResult, err := Collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Println(err)
		return err
	}

	if deleteResult.DeletedCount == 0 {
		return errors.New("Non-existent id")
	}

	return nil
}

/// GetAll ...
func (s *MongoStorage) GetAll() (map[int]Employee, error) {
	s.Lock()
	defer s.Unlock()

	var employee Employee
	var employees map[int]Employee

	cursor, err := Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println(err)

		defer cursor.Close(context.TODO())

		return employees, err
	}

	employees = make(map[int]Employee)

	for cursor.Next(context.TODO()) {
		if err := cursor.Decode(&employee); err != nil {
			log.Println(err)
			return employees, err
		}

		employees[employee.ID] = employee
	}

	return employees, nil
}

// Collection ...
var Collection *mongo.Collection

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

	Collection = client.Database("rest_api").Collection("rest_api")
}
