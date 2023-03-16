package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword ... Hash password of choice
// @Summary Hash password of choice
// @Description Hash password of choice
// @Description and returns the hashed password.
// @Tags Utils
// @Param password body string true "Password"
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// HashPassword ... Compares password with the hashed string
// @Summary Compares password with the hashed string
// @Description Compares password with the hashed string
// @Tags Utils
// @Param password body string true "Password"
// @Param hash body string true "Password hash"
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
