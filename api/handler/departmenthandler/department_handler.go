package departmenthandler

import (
	db "github.com/alaskastorm/rest-api/db/departmentdb"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DepartmentErrorResponse ...
type DepartmentErrorResponse struct {
	Message string `json:"message"`
}

// DepartmentHandler ...
type DepartmentHandler struct {
	storage db.DepartmentStorage
}

// NewDepartmentHandler ...
func NewDepartmentHandler(storage db.DepartmentStorage) *DepartmentHandler {
	return &DepartmentHandler{storage: storage}
}

// CreateDepartment ...
func (h *DepartmentHandler) CreateDepartment(c *gin.Context) {
	var department db.Department

	if err := c.BindJSON(&department); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartmentErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err := h.storage.Insert(&department)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartmentErrorResponse{
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": department.ID,
	})
}

// UpdateDepartment ...
func (h *DepartmentHandler) UpdateDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartmentErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var department db.Department

	if err := c.BindJSON(&department); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartmentErrorResponse{
			Message: err.Error(),
		})

		return
	}

	if err := h.storage.Update(id, &department); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartmentErrorResponse{
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// GetDepartment ...
func (h *DepartmentHandler) GetDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartmentErrorResponse{
			Message: err.Error(),
		})
		return
	}

	department, err := h.storage.Get(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartmentErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, department)

}

// DeleteDepartment ...
func (h *DepartmentHandler) DeleteDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartmentErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if err := h.storage.Delete(id); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartmentErrorResponse{
			Message: err.Error(),
		})

		return
	}

	c.String(http.StatusOK, "Department was deleted")
}

// GetAllDepartments ...
func (h *DepartmentHandler) GetAllDepartments(c *gin.Context) {
	departments, err := h.storage.GetAll()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartmentErrorResponse{
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, departments)
}

// AddEmployeeToDepartment ...
func (h *DepartmentHandler) AddEmployeeToDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, DepartmentErrorResponse{
			Message: err.Error(),
		})

		return
	}

	employeeID, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartmentErrorResponse{
			Message: err.Error(),
		})

		return
	}

	if err := h.storage.InsertEmployeeIntoDepartment(id, employeeID); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartmentErrorResponse{
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "employeedb was added",
	})
}
