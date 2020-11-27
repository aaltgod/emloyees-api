package employeedb

import (
	"github.com/alaskastorm/rest-api/db"

	"reflect"
	"testing"
)

type testEmployee struct {
	name         string
	employee     Employee
	wantError    error
	wantEmployee Employee
}

func Test_Employee_Insert(t *testing.T) {
	tests := []testEmployee{
		{
			name: "success",
			employee: Employee{
				Name:   "Leo",
				Sex:    "male",
				Age:    25,
				Salary: 5000,
			},
			wantError: nil,
		},
		{
			name: "success large integers",
			employee: Employee{
				Name:   "Lee",
				Sex:    "female",
				Age:    30,
				Salary: 5000000000001,
			},
			wantError: nil,
		},
	}

	db.ConnectToMongo()

	storage := NewEmployeeMongoStorage()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := storage.Insert(&tt.employee); got != tt.wantError {
				t.Errorf("storage.Insert() = %v, want %v", got, tt.wantError)
			}
		})
	}
}

func Test_Employee_Get(t *testing.T) {
	oneTest := testEmployee{
		name: "success",
		wantEmployee: Employee{
			ID:     1,
			Name:   "Leo",
			Sex:    "male",
			Age:    25,
			Salary: 5000,
		},
		employee: Employee{
			ID: 1,
		},
		wantError: nil,
	}

	db.ConnectToMongo()

	storage := NewEmployeeMongoStorage()

	t.Run(oneTest.name, func(t *testing.T) {
		gotEmployee, gotError := storage.Get(oneTest.employee.ID)
		if gotEmployee != oneTest.wantEmployee {
			t.Errorf("storage.Get() = %v, want %v", gotEmployee, oneTest.wantEmployee)
		}

		if gotError != oneTest.wantError {
			t.Errorf("storage.Get() = %v, want %v", gotError, oneTest.wantError)
		}
	})
}

func Test_Employee_Update(t *testing.T) {
	oneTest := testEmployee{
		name: "success",
		employee: Employee{
			ID:     1,
			Name:   "Leo",
			Sex:    "male",
			Age:    26,
			Salary: 5500,
		},
		wantError: nil,
	}

	db.ConnectToMongo()

	storage := NewEmployeeMongoStorage()

	t.Run(oneTest.name, func(t *testing.T) {
		if got := storage.Update(oneTest.employee.ID, &oneTest.employee); got != oneTest.wantError {
			t.Errorf("storage.Update() = %v, want %v", got, oneTest.wantError)
		}
	})
}

func Test_Employee_Delete(t *testing.T) {
	oneTest := testEmployee{
		name: "success",
		employee: Employee{
			ID: 1,
		},
		wantError: nil,
	}

	db.ConnectToMongo()

	storage := NewEmployeeMongoStorage()

	t.Run(oneTest.name, func(t *testing.T) {
		if got := storage.Delete(oneTest.employee.ID); got != oneTest.wantError {
			t.Errorf("storage.Delete() = %v, want %v", got, oneTest.wantError)
		}
	})
}

func Test_Employee_GetAll(t *testing.T) {
	oneTest := []testEmployee{
		{
			wantEmployee: Employee{
				ID:     1,
				Name:   "Leo",
				Sex:    "male",
				Age:    25,
				Salary: 5000,
			},
			wantError: nil,
		},
		{
			wantEmployee: Employee{
				ID:     2,
				Name:   "Lee",
				Sex:    "female",
				Age:    30,
				Salary: 5000000000001,
			},
			wantError: nil,
		},
	}

	wantEmployees := make(map[int]Employee)
	for _, e := range oneTest {
		wantEmployees[e.wantEmployee.ID] = e.wantEmployee
	}

	db.ConnectToMongo()

	storage := NewEmployeeMongoStorage()

	t.Run("success", func(t *testing.T) {
		gotEmployess, gotError := storage.GetAll()
		equal := reflect.DeepEqual(gotEmployess, wantEmployees)
		if !equal {
			t.Errorf("storage.GetAll() = %v, want %v", gotEmployess, wantEmployees)
		}
		if gotError != nil {
			t.Errorf("storage.GetAll() = %v, want %v", gotEmployess, wantEmployees)
		}
	})
}
