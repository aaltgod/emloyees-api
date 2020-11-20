package main

import (
	"context"
	"errors"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

// Employee ...
type Employee struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Sex             string `json:"sex"`
	Age             int    `json:"age"`
	Salary          int    `json:"salary"`
	DepartamentName string `json:"departament"`
}

// EmployeeStorage ...
type EmployeeStorage interface {
	Insert(e *Employee) error
	Get(id int) (Employee, error)
	Update(id int, e *Employee) error
	Delete(id int) error
	GetAll() (map[int]Employee, error)
}

// EmployeeMongoStorage ...
type EmployeeMongoStorage struct {
	counter int
	sync.Mutex
}

// NewEmployeeMongoStorage ...
func NewEmployeeMongoStorage() *EmployeeMongoStorage {
	return &EmployeeMongoStorage{
		counter: 1,
	}
}

// Insert ...
func (s *EmployeeMongoStorage) Insert(e *Employee) error {
	s.Lock()
	defer s.Unlock()

	e.ID = s.counter
	_, err := EmployeeCollection.InsertOne(context.TODO(), bson.M{
		"id":     *&e.ID,
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
func (s *EmployeeMongoStorage) Get(id int) (Employee, error) {
	s.Lock()
	defer s.Unlock()

	var employee Employee

	err := EmployeeCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&employee)
	if err != nil {
		log.Println(err)
		return employee, err
	}

	return employee, nil
}

// Update ...
func (s *EmployeeMongoStorage) Update(id int, e *Employee) error {
	s.Lock()
	defer s.Unlock()

	_, err := EmployeeCollection.UpdateOne(
		context.TODO(),
		bson.M{"id": id},
		bson.D{{"$set",
			bson.M{
				"id":     id,
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
func (s *EmployeeMongoStorage) Delete(id int) error {
	s.Lock()
	defer s.Unlock()

	deleteResult, err := EmployeeCollection.DeleteOne(context.TODO(), bson.M{"id": id})
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
func (s *EmployeeMongoStorage) GetAll() (map[int]Employee, error) {
	s.Lock()
	defer s.Unlock()

	var (
		employee  Employee
		employees map[int]Employee
	)

	cursor, err := EmployeeCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println(err)

		defer cursor.Close(context.TODO())

		return employees, err
	}

	employees = make(map[int]Employee)

	for cursor.Next(context.TODO()) {
		if err := cursor.Decode(&employee); err != nil {
			log.Println("CURSOR:", err)
			return employees, err
		}

		employees[employee.ID] = employee
	}

	return employees, nil
}
