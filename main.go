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

	router.GET("/employees", employeeHandler.GetAllEmployees)
	router.POST("/employee", employeeHandler.CreateEmployee)
	router.GET("/employee/:id", employeeHandler.GetEmployee)
	router.PUT("/employee/:id", employeeHandler.UpdateEmployee)
	router.DELETE("/employee/:id", employeeHandler.DeleteEmployee)
	router.GET("/departaments", departamentHandler.GetAllDepartaments)
	router.POST("/departament", departamentHandler.CreateDepartament)
	router.GET("/departament/:id", departamentHandler.GetDepartament)
	router.PUT("/departament/:id", departamentHandler.UpdateDepartament)
	router.DELETE("/departament/:id", departamentHandler.DeleteDepartament)

	router.Run(":5001")
}
