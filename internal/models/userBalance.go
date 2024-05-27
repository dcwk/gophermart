package models

type UserBalance struct {
	UserId     int64   `json:"user_id"`
	Accrual    float64 `json:"accrual"`
	Withdrawal float64 `json:"withdrawal"`
}

func NewUserBalance(userID int64) *UserBalance {
	return &UserBalance{
		UserId:     userID,
		Accrual:    0,
		Withdrawal: 0,
	}
}

func (userBalance *UserBalance) DoAccrual(value float64) {
	userBalance.Accrual += value
}

func (userBalance *UserBalance) DoWithdrawal(value float64) {
	userBalance.Accrual -= value
	userBalance.Withdrawal += value
}

func (userBalance *UserBalance) IsWithdrawPossible(value float64) bool {
	return userBalance.Accrual >= value
}
