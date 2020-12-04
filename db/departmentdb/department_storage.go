package departmentdb

import (
	"context"
	"errors"
	"github.com/alaskastorm/rest-api/db"

	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

// Department ...
type Department struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	EmployeesNumber int    `json:"employees_number"`
	EmployeesIDS    []int  `json:"employees_ids"`
}

// DepartmentStorage ...
type DepartmentStorage interface {
	Insert(d *Department) error
	Get(id int) (Department, error)
	Update(id int, d *Department) error
	Delete(id int) error
	GetAll() (map[int]Department, error)
	InsertEmployeeIntoDepartment(id, employeeID int) error
}

// DepartmentMongoStorage ...
type DepartmentMongoStorage struct {
	sync.Mutex
	counter int
}

// NewDepartmentMongoStorage ...
func NewDepartmentMongoStorage() *DepartmentMongoStorage {
	return &DepartmentMongoStorage{
		counter: 1,
	}
}

// Insert ...
func (s *DepartmentMongoStorage) Insert(d *Department) error {
	s.Lock()
	defer s.Unlock()

	d.ID = s.counter
	_, err := db.DepartmentCollection.InsertOne(context.TODO(), bson.M{
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
func (s *DepartmentMongoStorage) Get(id int) (Department, error) {
	s.Lock()
	defer s.Unlock()

	var department Department

	err := db.DepartmentCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&department)
	if err != nil {
		log.Println(err)
		return department, err
	}

	return department, nil
}

// Update ...
func (s *DepartmentMongoStorage) Update(id int, d *Department) error {
	s.Lock()
	defer s.Unlock()

	_, err := db.DepartmentCollection.UpdateOne(
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
func (s *DepartmentMongoStorage) Delete(id int) error {
	s.Lock()
	defer s.Unlock()

	deleteResult, err := db.DepartmentCollection.DeleteOne(context.TODO(), bson.M{"id": id})
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
func (s *DepartmentMongoStorage) GetAll() (map[int]Department, error) {
	s.Lock()
	defer s.Unlock()

	var department Department
	var departments map[int]Department

	cursor, err := db.DepartmentCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println(err)

		defer cursor.Close(context.TODO())

		return departments, err
	}

	departments = make(map[int]Department)

	for cursor.Next(context.TODO()) {
		if err := cursor.Decode(&department); err != nil {
			log.Println("CURSOR:", err)
			return departments, err
		}

		departments[department.ID] = department
	}

	return departments, nil
}

// InsertEmployeeIntoDepartment ...
func (s *DepartmentMongoStorage) InsertEmployeeIntoDepartment(id, employeeID int) error {
	s.Lock()
	defer s.Unlock()

	var department Department

	err := db.DepartmentCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&department)
	if err != nil {
		log.Println(err)
		return err
	}

	employeesIDS := department.EmployeesIDS
	employeesIDS = append(employeesIDS, employeeID)

	_, err = db.DepartmentCollection.UpdateOne(
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
