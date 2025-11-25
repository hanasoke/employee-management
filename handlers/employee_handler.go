package handlers

import (
	"employee-management/database"
	"employee-management/models"
	"employee-management/utils"
	"net/http"
	"time"

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
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	// Filter by department
	if department := c.Query("department"); department != "" {
		query = query.Where("department = ?", department)
	}

	// Search by name or NIK
	if search := c.Query("search"); search != "" {
		query = query.Where("name LIKE ? OR nik LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Find(&employees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  employees,
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

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// CreateEmployee - Create new employee
func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var req models.EmployeeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate email
	if !utils.IsValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	// Validate hire date
	if !utils.IsValidDate(req.HireDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hire date format. Use YYYY-MM-DD"})
		return
	}

	// Validate status if provided
	if req.Status != "" && !utils.IsValidStatus(req.Status) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	hireDate, _ := time.Parse("2006-01-02", req.HireDate)

	employee := models.Employee{
		NIK:        req.NIK,
		Name:       req.Name,
		Email:      req.Email,
		Position:   req.Position,
		Department: req.Department,
		Salary:     req.Salary,
		HireDate:   hireDate,
		Status:     req.Status,
	}

	if employee.Status == "" {
		employee.Status = "active"
	}

	// Check if NIK or Email already exists 
	var existingEmployee models.Employee 
	if err := h.DB.Where("nik = ? OR email = ?", employee.NIK, employee.Email).First(&existingEmployee).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NIK or Email already exists"})
		return 
	}

	if err := h.DB.Create(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Employee created successfully",
		"data": employee,
	})
}

// UpdateEmployee - Update employee by ID 
func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee
	if err := h.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return 
	}

	var req models.EmployeeRequest 
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	// Validate email 
	if !utils.IsValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
	}

	// Validate hire date 
	if !utils.IsValidDate(req.HireDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hire date format. Use YYYY-MM-DD"})
		return 
	}

	// Validate status if provided
	if req.Status != "" && !utils.IsValidStatus(req.Status) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return 
	}

	hireDate, _ := time.Parse("2006-01-02", req.HireDate)

	// check if NIK or Email already exists (excluding current employee)
	var existingEmployee models.Employee 
	if err := h.DB.Where("(nik = ? OR email = ?) AND id != ?", req.NIK, req.Email, id).First(&existingEmployee).Error; 
	err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NIK or Email already exists"})
		return 
	}

	// Update employee
    updates := models.Employee{
        NIK        : req.NIK,
        Name	   : req.Name,
        Email      : req.Email,
        Position   : req.Position,
        Department : req.Department,
        Salary     : req.Salary,
        HireDate   : hireDate,
        Status     : req.Status,
    }

	if err := h.DB.Model(&employee).Updates(updates).Error; 
	err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Employee updated successfully",
		"data": employee,
	})

}