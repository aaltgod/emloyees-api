package main

import (
	"github.com/alaskastorm/rest-api/api/handler/departamenthandler"
	"github.com/alaskastorm/rest-api/api/handler/employeehandler"
	"github.com/alaskastorm/rest-api/db"
	"github.com/alaskastorm/rest-api/db/departamentdb"
	"github.com/alaskastorm/rest-api/db/employeedb"
	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectToMongo()

	employeeStorage := employeedb.NewEmployeeMongoStorage()
	departamentStorage := departamentdb.NewDepartamentMongoStorage()
	employeeHandler := employeehandler.NewEmployeeHandler(employeeStorage)
	departamentHandler := departamenthandler.NewDepartamentHandler(departamentStorage)

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

	router.Run(":5245")
}
