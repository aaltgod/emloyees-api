package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// EmployeeErrorResponse ...
type EmployeeErrorResponse struct {
	Message string `json:"message"`
}

// EmployeeHandler ...
type EmployeeHandler struct {
	storage EmployeeStorage
}

// NewEmployeeHandler ...
func NewEmployeeHandler(storage EmployeeStorage) *EmployeeHandler {
	return &EmployeeHandler{storage: storage}
}

// CreateEmployee ...
func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var employee Employee

	if err := c.BindJSON(&employee); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, EmployeeErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err := h.storage.Insert(&employee)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, EmployeeErrorResponse{
			Message: "Ooops, Something went wrong",
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": employee.ID,
	})
}

// UpdateEmployee ...
func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, EmployeeErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var employee Employee

	if err := c.BindJSON(&employee); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, EmployeeErrorResponse{
			Message: "Ooops, Something went wrong",
		})

		return
	}

	if err := h.storage.Update(id, &employee); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, EmployeeErrorResponse{
			Message: "Ooops, Something went wrong",
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// GetEmployee ...
func (h *EmployeeHandler) GetEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, EmployeeErrorResponse{
			Message: err.Error(),
		})
		return
	}

	employee, err := h.storage.Get(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, EmployeeErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// DeleteEmployee ...
func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, EmployeeErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if err := h.storage.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, EmployeeErrorResponse{
			Message: "Ooops, Something went wrong",
		})

		return
	}

	c.String(http.StatusOK, "Employee was deleted")
}

// GetAllEmployees ...
func (h *EmployeeHandler) GetAllEmployees(c *gin.Context) {
	employees, err := h.storage.GetAll()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, EmployeeErrorResponse{
			Message: "Ooops, Something went wrong",
		})

		return
	}

	c.JSON(http.StatusOK, employees)
}
