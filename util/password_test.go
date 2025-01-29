package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "securepassword123"

	// Test successful hashing
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err, "HashPassword should not return an error")
	assert.NotEmpty(t, hashedPassword, "Hashed password should not be empty")

	// Test that the hashed password is different from the original
	assert.NotEqual(t, password, hashedPassword, "Hashed password should not match the plain password")

	// Test CheckPassword with the correct password
	err = CheckPassword(password, hashedPassword)
	assert.NoError(t, err, "CheckPassword should not return an error for the correct password")

	// Test CheckPassword with an incorrect password
	err = CheckPassword("wrongpassword", hashedPassword)
	assert.Error(t, err, "CheckPassword should return an error for an incorrect password")
}

func TestCheckPassword(t *testing.T) {
	// Hash a password for testing
	password := "testpassword"
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err, "HashPassword should not return an error")

	// Check the hashed password against the correct plain password
	err = CheckPassword(password, hashedPassword)
	assert.NoError(t, err, "CheckPassword should not return an error for the correct password")

	// Check the hashed password against an incorrect plain password
	err = CheckPassword("invalidpassword", hashedPassword)
	assert.Error(t, err, "CheckPassword should return an error for an incorrect password")
}
