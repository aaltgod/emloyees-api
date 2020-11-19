package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DepartamentErrorResponse ...
type DepartamentErrorResponse struct {
	Message string `json:"message"`
}

// DepartamentHandler ...
type DepartamentHandler struct {
	storage DepartamentStorage
}

// NewDepartamentHandler ...
func NewDepartamentHandler(storage DepartamentStorage) *DepartamentHandler {
	return &DepartamentHandler{storage: storage}
}

// CreateDepartament ...
func (h *DepartamentHandler) CreateDepartament(c *gin.Context) {
	var departament Departament

	if err := c.BindJSON(&departament); err != nil {
		fmt.Printf("failed to bind an departament: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err := h.storage.Insert(&departament)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: "Ooops, Something went wrong",
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": departament.ID,
	})
}

// UpdateDepartament ...
func (h *DepartamentHandler) UpdateDepartament(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var departament Departament

	if err := c.BindJSON(&departament); err != nil {
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: "Ooops, Something went wrong",
		})

		return
	}

	if err := h.storage.Update(id, &departament); err != nil {
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: "Ooops, Something went wrong",
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// GetDepartament ...
func (h *DepartamentHandler) GetDepartament(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: err.Error(),
		})
		return
	}

	departament, err := h.storage.Get(id)
	if err != nil {
		fmt.Printf("failed to get departament: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, departament)
}

// DeleteDepartament ...
func (h *DepartamentHandler) DeleteDepartament(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failde to convert id to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if err := h.storage.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: "Ooops, Something went wrong",
		})

		return
	}

	c.String(http.StatusOK, "Departament was deleted")
}

// GetAllDepartaments ...
func (h *DepartamentHandler) GetAllDepartaments(c *gin.Context) {
	departaments, err := h.storage.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: "Ooops, Something went wrong",
		})

		return
	}
	fmt.Println(departaments)

	c.JSON(http.StatusOK, departaments)
}
