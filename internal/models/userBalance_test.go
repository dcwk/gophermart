package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserBalance_DoAccrual(t *testing.T) {
	tests := []struct {
		Name    string
		UserID  int64
		Accrual float64
		Want    float64
	}{
		{
			Name:    "Test can check account balance",
			UserID:  1,
			Accrual: 100,
			Want:    100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			userBalance := NewUserBalance(tt.UserID)

			userBalance.DoAccrual(tt.Accrual)

			assert.Equal(t, userBalance.Accrual, tt.Want)
		})
	}
}

func TestUserBalance_DoWithdrawal(t *testing.T) {
	tests := []struct {
		Name       string
		UserID     int64
		Accrual    float64
		Withdrawal float64
		Want       float64
	}{
		{
			Name:       "Test can add withdrawal",
			UserID:     1,
			Accrual:    500,
			Withdrawal: 100,
			Want:       400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			userBalance := NewUserBalance(tt.UserID)
			userBalance.DoAccrual(tt.Accrual)

			userBalance.DoWithdrawal(tt.Withdrawal)

			assert.Equal(t, userBalance.Accrual, tt.Want)
		})
	}
}

func TestUserBalance_IsWithdrawPossible(t *testing.T) {
	tests := []struct {
		Name       string
		UserID     int64
		Accrual    float64
		Withdrawal float64
		Want       bool
	}{
		{
			Name:       "Test can check account withdraw is possible",
			UserID:     1,
			Accrual:    500,
			Withdrawal: 100,
			Want:       true,
		},
		{
			Name:       "Test fail check account withdraw is possible",
			UserID:     1,
			Accrual:    500,
			Withdrawal: 900,
			Want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			userBalance := NewUserBalance(tt.UserID)
			userBalance.DoAccrual(tt.Accrual)

			res := userBalance.IsWithdrawPossible(tt.Withdrawal)

			assert.Equal(t, res, tt.Want)
		})
	}
}
