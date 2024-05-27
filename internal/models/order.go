package models

import (
	"strconv"
	"time"
)

type Order struct {
	ID        int64     `json:id`
	UserID    int64     `json:user_id`
	Number    string    `json:number`
	Status    string    `json:status`
	Accrual   float64   `json:accepted`
	CreatedAt time.Time `json:created_at`
}

func NewOrder(userID int64, number string) *Order {
	return &Order{
		UserID:    userID,
		Number:    number,
		CreatedAt: time.Now(),
	}
}

func (o *Order) IsValid() bool {
	number, err := strconv.Atoi(o.Number)
	if err != nil {
		return false
	}

	return (number%10+checksum(number/10))%10 == 0
}

func checksum(number int) int {
	var luhn int

	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 { // even
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		luhn += cur
		number = number / 10
	}
	return luhn % 10
}
