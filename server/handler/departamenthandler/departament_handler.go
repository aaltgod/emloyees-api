package departamenthandler

import (
	db "github.com/alaskastorm/rest-api/db/departamentdb"
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
	storage db.DepartamentStorage
}

// NewDepartamentHandler ...
func NewDepartamentHandler(storage db.DepartamentStorage) *DepartamentHandler {
	return &DepartamentHandler{storage: storage}
}

// CreateDepartament ...
func (h *DepartamentHandler) CreateDepartament(c *gin.Context) {
	var departament db.Departament

	if err := c.BindJSON(&departament); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err := h.storage.Insert(&departament)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: err.Error(),
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
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var departament db.Departament

	if err := c.BindJSON(&departament); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: err.Error(),
		})

		return
	}

	if err := h.storage.Update(id, &departament); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: err.Error(),
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
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: err.Error(),
		})
		return
	}

	departament, err := h.storage.Get(id)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if err := h.storage.Delete(id); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: err.Error(),
		})

		return
	}

	c.String(http.StatusOK, "Departament was deleted")
}

// GetAllDepartaments ...
func (h *DepartamentHandler) GetAllDepartaments(c *gin.Context) {
	departaments, err := h.storage.GetAll()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, departaments)
}

// AddEmployeeToDepartament ...
func (h *DepartamentHandler) AddEmployeeToDepartament(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: err.Error(),
		})

		return
	}

	employeeID, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: err.Error(),
		})

		return
	}

	if err := h.storage.InsertEmployeeIntoDepartament(id, employeeID); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, DepartamentErrorResponse{
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "employeedb was added",
	})
}
