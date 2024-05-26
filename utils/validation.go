package utils

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func IsValidEmail(email string) bool {
	// Define a regex pattern for a basic email validation
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	// Check if the email matches the regex pattern
	return emailRegex.MatchString(email)
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
