package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccrual_UpdateStatus(t *testing.T) {
	tests := []struct {
		Name    string
		OrderID int64
		Status  string
		Value   float64
		Want    *Accrual
	}{
		{
			Name:    "Test can update status",
			OrderID: 1,
			Status:  Invalid,
			Value:   0,
			Want: &Accrual{
				OrderID: 1,
				Status:  Invalid,
				Value:   0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			accrual := NewAccrual(tt.OrderID)
			assert.Equal(t, accrual.Status, Processing)
			assert.Equal(t, accrual.Value, float64(0))

			accrual.UpdateStatus(tt.Status, tt.Value)

			assert.Equal(t, tt.Want.Status, accrual.Status)
			assert.Equal(t, tt.Want.Value, accrual.Value)
		})
	}
}
