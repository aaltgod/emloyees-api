package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ConnectToMongo()

	employeeStorage := NewEmployeeMongoStorage()
	departamentStorage := NewDepartamentMongoStorage()
	employeeHandler := NewEmployeeHandler(employeeStorage)
	departamentHandler := NewDepartamentHandler(departamentStorage)

	router := gin.Default()

	router.GET("/api/employees", employeeHandler.GetAllEmployees)
	router.POST("/api/employee", employeeHandler.CreateEmployee)
	router.GET("/api/employee/:id", employeeHandler.GetEmployee)
	router.PUT("/api/employee/:id", employeeHandler.UpdateEmployee)
	router.DELETE("/api/employee/:id", employeeHandler.DeleteEmployee)
	router.GET("/api/departaments", departamentHandler.GetAllDepartaments)
	router.POST("/api/departament", departamentHandler.CreateDepartament)
	router.GET("/api/departament/:id", departamentHandler.GetDepartament)
	router.PUT("/api/departament/:id", departamentHandler.UpdateDepartament)
	router.DELETE("/api/departament/:id", departamentHandler.DeleteDepartament)
	router.PUT("/api/departament/:id/:employee_id", departamentHandler.AddEmployeeToDepartament)

	router.Run(":5001")
}
