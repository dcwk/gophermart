package models

type UserBalance struct {
	UserId     int64   `json:"user_id"`
	Accrual    float64 `json:"accrual"`
	Withdrawal float64 `json:"withdrawal"`
}

func NewUserBalance(userId int64) *UserBalance {
	return &UserBalance{
		UserId:     userId,
		Accrual:    0,
		Withdrawal: 0,
	}
}
