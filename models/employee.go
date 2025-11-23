package models

import (
	"time"
)

type Employee struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	NIK        string    `json:"nik" gorm:"unique;not null"`
	Name       string    `json:"name" gorm:"not null"`
	Email      string    `json:"email" gorm:"unique;not null"`
	Position   string    `json:"position" gorm:"not null"`
	Department string    `json:"department" gorm:"not null"`
	Salary     float64   `json:"salary" gorm:"not null"`
	HireDate   time.Time `json:"hire_date"`
	Status     string    `json:"status" gorm:"default:'active'"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type EmployeeRequest struct {
	NIK       string  `json:"nik" binding:"required"`
    Name      string  `json:"name" binding:"required"`
    Email     string  `json:"email" binding:"required,email"`
    Position  string  `json:"position" binding:"required"`
    Department string `json:"department" binding:"required"`
    Salary    float64 `json:"salary" binding:"required,min=0"`
    HireDate  string  `json:"hire_date" binding:"required"`
    Status    string  `json:"status"`
}