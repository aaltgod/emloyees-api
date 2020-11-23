package main

import (
	"reflect"
	"testing"
)

type testDepartament struct {
	name            string
	departament     Departament
	wantError       error
	wantDepartament Departament
}

func Test_Departament_Insert(t *testing.T) {
	tests := []testDepartament{
		{
			name: "success №1",
			departament: Departament{
				Name: "EvilRabbits",
			},
			wantError: nil,
		},
		{
			name: "success №2",
			departament: Departament{
				Name: "HappyRacoons",
			},
			wantError: nil,
		},
	}

	ConnectToMongo()

	storage := NewDepartamentMongoStorage()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := storage.Insert(&tt.departament); got != tt.wantError {
				t.Errorf("storage.Insert() = %v, want %v", got, tt.wantError)
			}
		})
	}
}

func Test_Departament_Get(t *testing.T) {
	oneTest := testDepartament{
		name: "success",
		wantDepartament: Departament{
			ID:              1,
			Name:            "EvilRabbits",
			EmployeesNumber: 0,
			EmployeesIDS:    []int{},
		},
		departament: Departament{
			ID: 1,
		},
		wantError: nil,
	}

	ConnectToMongo()

	storage := NewDepartamentMongoStorage()

	t.Run(oneTest.name, func(t *testing.T) {
		gotDepartament, gotError := storage.Get(oneTest.departament.ID)
		if gotDepartament.Name != oneTest.wantDepartament.Name {
			t.Errorf(
				"storage.Get() = %v, want %v",
				gotDepartament.Name,
				oneTest.wantDepartament.Name,
			)
		}

		if gotDepartament.ID != oneTest.wantDepartament.ID {
			t.Errorf(
				"storage.Get() = %v, want %v",
				gotDepartament.ID,
				oneTest.wantDepartament.ID,
			)
		}

		if gotDepartament.EmployeesNumber != oneTest.wantDepartament.EmployeesNumber {
			t.Errorf(
				"storage.Get() = %v, want %v",
				gotDepartament.EmployeesNumber,
				oneTest.wantDepartament.EmployeesNumber,
			)
		}

		equal := reflect.DeepEqual(gotDepartament.EmployeesIDS, oneTest.wantDepartament.EmployeesIDS)
		if !equal {
			t.Errorf(
				"storage.Get() = %v, want %v",
				gotDepartament.EmployeesIDS,
				oneTest.wantDepartament.EmployeesIDS)
		}

		if gotError != oneTest.wantError {
			t.Errorf("storage.Get() = %v, want %v", gotError, oneTest.wantError)
		}
	})
}

func Test_Departament_Update(t *testing.T) {
	oneTest := testDepartament{
		name: "success",
		departament: Departament{
			ID:   1,
			Name: "TiredRabbits",
		},
		wantError: nil,
	}

	ConnectToMongo()

	storage := NewDepartamentMongoStorage()

	t.Run(oneTest.name, func(t *testing.T) {
		if got := storage.Update(oneTest.departament.ID, &oneTest.departament); got != oneTest.wantError {
			t.Errorf("storage.Update() = %v, want %v", got, oneTest.wantError)
		}
	})
}

func Test_Departament_Delete(t *testing.T) {
	oneTest := testDepartament{
		name: "success",
		departament: Departament{
			ID: 1,
		},
		wantError: nil,
	}

	ConnectToMongo()

	storage := NewDepartamentMongoStorage()

	t.Run(oneTest.name, func(t *testing.T) {
		if got := storage.Delete(oneTest.departament.ID); got != oneTest.wantError {
			t.Errorf("storage.Delete() = %v, want %v", got, oneTest.wantError)
		}
	})
}

func Test_Departament_GetAll(t *testing.T) {

	oneTest := []testDepartament{
		{
			wantDepartament: Departament{
				ID:              1,
				Name:            "EvilRabbits",
				EmployeesNumber: 0,
				EmployeesIDS:    []int{},
			},
			wantError: nil,
		},
		{
			wantDepartament: Departament{
				ID:              2,
				Name:            "HappyRacoons",
				EmployeesNumber: 0,
				EmployeesIDS:    []int{},
			},
			wantError: nil,
		},
	}

	wantDepartaments := make(map[int]Departament)
	for _, e := range oneTest {
		wantDepartaments[e.wantDepartament.ID] = e.wantDepartament
	}

	ConnectToMongo()

	storage := NewDepartamentMongoStorage()

	t.Run("success", func(t *testing.T) {
		gotDepartaments, gotError := storage.GetAll()
		equal := reflect.DeepEqual(gotDepartaments, wantDepartaments)
		if !equal {
			t.Errorf("storage.GetAll() = %v, want %v", gotDepartaments, wantDepartaments)
		}
		if gotError != nil {
			t.Errorf("storage.GetAll() = %v, want %v", gotDepartaments, wantDepartaments)
		}
	})
}

func Test_Departament_InsertEmployeeIntoDepartament(t *testing.T) {
	oneTest := testDepartament{
		name: "success",
		departament: Departament{
			ID: 1,
		},
		wantError: nil,
	}
	employeeID := 1

	ConnectToMongo()

	storage := NewDepartamentMongoStorage()

	t.Run(oneTest.name, func(t *testing.T) {
		gotError := storage.InsertEmployeeIntoDepartament(oneTest.departament.ID, employeeID)
		if gotError != oneTest.wantError {
			t.Errorf(
				"storage.InsertEmployeeIntoDepartament() = %v, want %v",
				gotError, oneTest.wantError)
		}
	})
}
