package main

import (
	"context"
	"errors"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

// Departament ...
type Departament struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	EmployeesNumber int    `json:"employees_number"`
	EmployeesIDS    []int  `json:"employees_ids"`
}

// DepartamentStorage ...
type DepartamentStorage interface {
	Insert(d *Departament) error
	Get(id int) (Departament, error)
	Update(id int, d *Departament) error
	Delete(id int) error
	GetAll() (map[int]Departament, error)
	InsertEmployeeIntoDepartament(id, employeeID int) error
}

// DepartamentMongoStorage ...
type DepartamentMongoStorage struct {
	sync.Mutex
	counter int
}

// NewDepartamentMongoStorage ...
func NewDepartamentMongoStorage() *DepartamentMongoStorage {
	return &DepartamentMongoStorage{
		counter: 1,
	}
}

// Insert ...
func (s *DepartamentMongoStorage) Insert(d *Departament) error {
	s.Lock()
	defer s.Unlock()

	d.ID = s.counter
	_, err := DepartamentCollection.InsertOne(context.TODO(), bson.M{
		"id":              *&d.ID,
		"name":            *&d.Name,
		"employeesNumber": 0,
		"employeesIDS":    []int{},
	})
	if err != nil {
		log.Println(err)
		return err
	}

	s.counter++

	return nil
}

// Get ...
func (s *DepartamentMongoStorage) Get(id int) (Departament, error) {
	s.Lock()
	defer s.Unlock()

	var departament Departament

	err := DepartamentCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&departament)
	if err != nil {
		log.Println(err)
		return departament, err
	}

	return departament, nil
}

// Update ...
func (s *DepartamentMongoStorage) Update(id int, d *Departament) error {
	s.Lock()
	defer s.Unlock()

	_, err := DepartamentCollection.UpdateOne(
		context.TODO(),
		bson.M{"id": id},
		bson.D{{"$set", bson.M{"name": *&d.Name}}})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// Delete ...
func (s *DepartamentMongoStorage) Delete(id int) error {
	s.Lock()
	defer s.Unlock()

	deleteResult, err := DepartamentCollection.DeleteOne(context.TODO(), bson.M{"id": id})
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
func (s *DepartamentMongoStorage) GetAll() (map[int]Departament, error) {
	s.Lock()
	defer s.Unlock()

	var departament Departament
	var departaments map[int]Departament

	cursor, err := DepartamentCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println(err)

		defer cursor.Close(context.TODO())

		return departaments, err
	}

	departaments = make(map[int]Departament)

	for cursor.Next(context.TODO()) {
		if err := cursor.Decode(&departament); err != nil {
			log.Println("CURSOR:", err)
			return departaments, err
		}

		departaments[departament.ID] = departament
	}

	return departaments, nil
}

// InsertEmployeeIntoDepartament ...
func (s *DepartamentMongoStorage) InsertEmployeeIntoDepartament(id, employeeID int) error {
	s.Lock()
	defer s.Unlock()

	var departament Departament

	err := DepartamentCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&departament)
	if err != nil {
		log.Println(err)
		return err
	}

	employeesIDS := departament.EmployeesIDS
	employeesIDS = append(employeesIDS, employeeID)

	_, err = DepartamentCollection.UpdateOne(
		context.TODO(),
		bson.M{"id": id},
		bson.D{{"$set", bson.M{
			"employeesNumber": len(employeesIDS),
			"employeesIDS":    employeesIDS,
		}}},
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}
