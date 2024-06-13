package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCanHashUserPassword(t *testing.T) {
	passwordString := "testPassword"
	user := User{
		ID:       1,
		Login:    "testLogin",
		Password: passwordString,
	}

	t.Run("Test can hash user password", func(t *testing.T) {
		assert.NoError(t, user.HashPassword())

		assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordString)))
	})
}

func TestCanVerifyPassword(t *testing.T) {
	passwordString := "testPassword"
	user := User{
		ID:       1,
		Login:    "testLogin",
		Password: passwordString,
	}

	t.Run("Test can verify user password", func(t *testing.T) {
		assert.NoError(t, user.HashPassword())

		assert.True(t, user.VerifyPassword(passwordString))
	})
}
