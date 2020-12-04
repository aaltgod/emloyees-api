package departmentdb

import (
	"github.com/alaskastorm/rest-api/db"
	"reflect"
	"testing"
)

type testDepartment struct {
	name           string
	department     Department
	wantError      error
	wantDepartment Department
}

func Test_Department_Insert(t *testing.T) {
	tests := []testDepartment{
		{
			name: "success №1",
			department: Department{
				Name: "EvilRabbits",
			},
			wantError: nil,
		},
		{
			name: "success №2",
			department: Department{
				Name: "HappyRacoons",
			},
			wantError: nil,
		},
	}

	db.ConnectToMongo()

	storage := NewDepartmentMongoStorage()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := storage.Insert(&tt.department); got != tt.wantError {
				t.Errorf("storage.Insert() = %v, want %v", got, tt.wantError)
			}
		})
	}
}

func Test_Department_Get(t *testing.T) {
	oneTest := testDepartment{
		name: "success",
		wantDepartment: Department{
			ID:              1,
			Name:            "EvilRabbits",
			EmployeesNumber: 0,
			EmployeesIDS:    []int{},
		},
		department: Department{
			ID: 1,
		},
		wantError: nil,
	}

	db.ConnectToMongo()

	storage := NewDepartmentMongoStorage()

	t.Run(oneTest.name, func(t *testing.T) {
		gotDepartment, gotError := storage.Get(oneTest.department.ID)
		if gotDepartment.Name != oneTest.wantDepartment.Name {
			t.Errorf(
				"storage.Get() = %v, want %v",
				gotDepartment.Name,
				oneTest.wantDepartment.Name,
			)
		}

		if gotDepartment.ID != oneTest.wantDepartment.ID {
			t.Errorf(
				"storage.Get() = %v, want %v",
				gotDepartment.ID,
				oneTest.wantDepartment.ID,
			)
		}

		if gotDepartment.EmployeesNumber != oneTest.wantDepartment.EmployeesNumber {
			t.Errorf(
				"storage.Get() = %v, want %v",
				gotDepartment.EmployeesNumber,
				oneTest.wantDepartment.EmployeesNumber,
			)
		}

		equal := reflect.DeepEqual(gotDepartment.EmployeesIDS, oneTest.wantDepartment.EmployeesIDS)
		if !equal {
			t.Errorf(
				"storage.Get() = %v, want %v",
				gotDepartment.EmployeesIDS,
				oneTest.wantDepartment.EmployeesIDS)
		}

		if gotError != oneTest.wantError {
			t.Errorf("storage.Get() = %v, want %v", gotError, oneTest.wantError)
		}
	})
}

func Test_Department_Update(t *testing.T) {
	oneTest := testDepartment{
		name: "success",
		department: Department{
			ID:   1,
			Name: "TiredRabbits",
		},
		wantError: nil,
	}

	db.ConnectToMongo()

	storage := NewDepartmentMongoStorage()

	t.Run(oneTest.name, func(t *testing.T) {
		if got := storage.Update(oneTest.department.ID, &oneTest.department); got != oneTest.wantError {
			t.Errorf("storage.Update() = %v, want %v", got, oneTest.wantError)
		}
	})
}

func Test_Department_Delete(t *testing.T) {
	oneTest := testDepartment{
		name: "success",
		department: Department{
			ID: 1,
		},
		wantError: nil,
	}

	db.ConnectToMongo()

	storage := NewDepartmentMongoStorage()

	t.Run(oneTest.name, func(t *testing.T) {
		if got := storage.Delete(oneTest.department.ID); got != oneTest.wantError {
			t.Errorf("storage.Delete() = %v, want %v", got, oneTest.wantError)
		}
	})
}

func Test_Department_GetAll(t *testing.T) {

	oneTest := []testDepartment{
		{
			wantDepartment: Department{
				ID:              1,
				Name:            "EvilRabbits",
				EmployeesNumber: 0,
				EmployeesIDS:    []int{},
			},
			wantError: nil,
		},
		{
			wantDepartment: Department{
				ID:              2,
				Name:            "HappyRacoons",
				EmployeesNumber: 0,
				EmployeesIDS:    []int{},
			},
			wantError: nil,
		},
	}

	wantDepartments := make(map[int]Department)
	for _, e := range oneTest {
		wantDepartments[e.wantDepartment.ID] = e.wantDepartment
	}

	db.ConnectToMongo()

	storage := NewDepartmentMongoStorage()

	t.Run("success", func(t *testing.T) {
		gotDepartments, gotError := storage.GetAll()
		equal := reflect.DeepEqual(gotDepartments, wantDepartments)
		if !equal {
			t.Errorf("storage.GetAll() = %v, want %v", gotDepartments, wantDepartments)
		}
		if gotError != nil {
			t.Errorf("storage.GetAll() = %v, want %v", gotDepartments, wantDepartments)
		}
	})
}

func Test_Department_InsertEmployeeIntoDepartment(t *testing.T) {
	oneTest := testDepartment{
		name: "success",
		department: Department{
			ID: 1,
		},
		wantError: nil,
	}
	employeeID := 1

	db.ConnectToMongo()

	storage := NewDepartmentMongoStorage()

	t.Run(oneTest.name, func(t *testing.T) {
		gotError := storage.InsertEmployeeIntoDepartment(oneTest.department.ID, employeeID)
		if gotError != oneTest.wantError {
			t.Errorf(
				"storage.InsertEmployeeIntoDepartment() = %v, want %v",
				gotError, oneTest.wantError)
		}
	})
}
