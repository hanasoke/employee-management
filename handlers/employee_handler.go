package handlers

import (
	"employee-management/database"
	"employee-management/models"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	DB *gorm.DB 
}

func NewEmployeeHandler() *EmployeeHandler {
	return &EmployeeHandler{DB: database.DB}
}

// GetEmployees - Get all employees with optional filtering 
func (h *EmployeeHandler) GetEmployees(c *gin.Context) {
	var employees []models.Employee

	query := h.DB.Model(&models.Employee{})

	// Filter by status 
	if status := c.Query("status"); 
	status != "" {
		query = query.Where("status = ?", status)
	}

	// Filter by department 
	if department := c.Query("department"); 
	department != "" {
		query = query.Where("department = ?", department)
	}

	// Search by name or NIK 
	if search := c.Query("search");
	search != "" {
		query = query.Where("name LIKE ? OR nik LIKE ?", "%"+search+"%", "%"+search+"%")
	}


	if err := query.Find(&employees).Error; 
	err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"data": employees,
		"count": len(employees),
	})

}

// GetEmployee - Get employee by ID 
func (h *EmployeeHandler) GetEmployee(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee 
	if err := h.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"data":employee})
}