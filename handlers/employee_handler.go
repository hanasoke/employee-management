package handlers

import (
	"employee-management/database"
	"employee-management/models"

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

	
	

}
