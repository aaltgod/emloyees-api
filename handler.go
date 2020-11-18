package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ErrorResponse ...
type ErrorResponse struct {
	Message string `json:"message"`
}

// Handler ...
type Handler struct {
	storage Storage
}

// NewHandler ...
func NewHandler(storage Storage) *Handler {
	return &Handler{storage: storage}
}

// CreateEmployee ...
func (h *Handler) CreateEmployee(c *gin.Context) {
	var employee Employee

	if err := c.BindJSON(&employee); err != nil {
		fmt.Printf("failed to bind an employee: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err := h.storage.Insert(&employee)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Ooops, Something went wrong",
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": employee.ID,
	})
}

// UpdateEmployee ...
func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var employee Employee

	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Ooops, Something went wrong",
		})

		return
	}

	if err := h.storage.Update(id, &employee); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Ooops, Something went wrong",
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// GetEmployee ...
func (h *Handler) GetEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	employee, err := h.storage.Get(id)
	if err != nil {
		fmt.Printf("failed to get employee: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// DeleteEmployee ...
func (h *Handler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failde to convert id to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if err := h.storage.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Ooops, Something went wrong",
		})

		return
	}

	c.String(http.StatusOK, "Employee was deleted")
}

// GetAllEmployees ...
func (h *Handler) GetAllEmployees(c *gin.Context) {
	employees, err := h.storage.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Ooops, Something went wrong",
		})

		return
	}
	fmt.Println(employees)

	c.JSON(http.StatusOK, employees)
}
