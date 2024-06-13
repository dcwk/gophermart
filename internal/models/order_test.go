package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrder_IsValid(t *testing.T) {
	tests := []struct {
		Name        string
		UserID      int64
		OrderNumber string
		Want        bool
	}{
		{
			Name:        "Test can validate order",
			UserID:      1,
			OrderNumber: "110135840886155",
			Want:        true,
		},
		{
			Name:        "Test fail validate order",
			UserID:      1,
			OrderNumber: "123",
			Want:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			order := NewOrder(tt.UserID, tt.OrderNumber)

			res := order.IsValid()

			assert.Equal(t, tt.Want, res)
		})
	}
}
