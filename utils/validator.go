package utils

import (
	"regexp"
	"time"
)

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func IsValidDate(dataStr string) bool {
	_, err := time.Parse("2006-01-02", dataStr)
	return err == nil
}

func IsValidStatus(status string) bool {
	validStatuses := []string{"active", "inactive", "resigned", "terminated"}
	for _, s := range validStatuses {
		if s == status {
			return true
		}
	}
	return false
}
