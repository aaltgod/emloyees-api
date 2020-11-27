package main

import (
	"github.com/alaskastorm/rest-api/db"
	"github.com/alaskastorm/rest-api/db/departamentdb"
	"github.com/alaskastorm/rest-api/db/employeedb"
	"github.com/alaskastorm/rest-api/server/handler/employeehandler"
	"github.com/alaskastorm/rest-api/server/handler/departamenthandler"
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
	router.POST("/api/employeedb", employeeHandler.CreateEmployee)
	router.GET("/api/employeedb/:id", employeeHandler.GetEmployee)
	router.PUT("/api/employeedb/:id", employeeHandler.UpdateEmployee)
	router.DELETE("/api/employeedb/:id", employeeHandler.DeleteEmployee)
	router.GET("/api/departaments", departamentHandler.GetAllDepartaments)
	router.POST("/api/departamentdb", departamentHandler.CreateDepartament)
	router.GET("/api/departamentdb/:id", departamentHandler.GetDepartament)
	router.PUT("/api/departamentdb/:id", departamentHandler.UpdateDepartament)
	router.DELETE("/api/departamentdb/:id", departamentHandler.DeleteDepartament)
	router.PUT("/api/departamentdb/:id/:employee_id", departamentHandler.AddEmployeeToDepartament)

	router.Run(":5001")
}
