package models

import "time"

type Order struct {
	ID        int64     `json:id`
	UserID    int64     `json:user_id`
	Number    string    `json:number`
	CreatedAt time.Time `json:created_at`
}
