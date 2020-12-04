package main

import (
	"github.com/alaskastorm/rest-api/api/handler/departmenthandler"
	"github.com/alaskastorm/rest-api/api/handler/employeehandler"
	"github.com/alaskastorm/rest-api/db"
	"github.com/alaskastorm/rest-api/db/departmentdb"
	"github.com/alaskastorm/rest-api/db/employeedb"
	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectToMongo()

	employeeStorage := employeedb.NewEmployeeMongoStorage()
	departmentStorage := departmentdb.NewDepartmentMongoStorage()
	employeeHandler := employeehandler.NewEmployeeHandler(employeeStorage)
	departmentHandler := departmenthandler.NewDepartmentHandler(departmentStorage)

	router := gin.Default()

	router.GET("/api/employees", employeeHandler.GetAllEmployees)
	router.POST("/api/employee", employeeHandler.CreateEmployee)
	router.GET("/api/employee/:id", employeeHandler.GetEmployee)
	router.PUT("/api/employee/:id", employeeHandler.UpdateEmployee)
	router.DELETE("/api/employee/:id", employeeHandler.DeleteEmployee)
	router.GET("/api/departments", departmentHandler.GetAllDepartments)
	router.POST("/api/department", departmentHandler.CreateDepartment)
	router.GET("/api/department/:id", departmentHandler.GetDepartment)
	router.PUT("/api/department/:id", departmentHandler.UpdateDepartment)
	router.DELETE("/api/department/:id", departmentHandler.DeleteDepartment)
	router.PUT("/api/department/:id/:employee_id", departmentHandler.AddEmployeeToDepartment)

	router.Run(":5245")
}
