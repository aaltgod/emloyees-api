package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ConnectToMongo()
	mongoStorage := NewMongoStorage()
	handler := NewHandler(mongoStorage)

	router := gin.Default()

	router.GET("/", handler.GetAllEmployees)
	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	router.Run(":5001")
}
