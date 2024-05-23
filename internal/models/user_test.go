package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

// TODO: разобраться почему не проходят тесты
func TestCanHashUserPassword(t *testing.T) {
	passwordString := "testPassword"
	user := User{
		ID:       1,
		Login:    "testLogin",
		Password: passwordString,
	}

	t.Run("Test can hash user password", func(t *testing.T) {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.DefaultCost)

		err = user.HashPassword()
		assert.NoError(t, err)

		assert.Equal(t, string(hashedPassword), user.Password)
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
		res, err := user.VerifyPassword(passwordString)

		assert.NoError(t, err)
		assert.True(t, res)
	})
}
