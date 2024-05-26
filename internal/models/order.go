package models

import "time"

type Order struct {
	ID        int       `json:id`
	UserID    int       `json:user_id`
	Number    string    `json:number`
	CreatedAt time.Time `json:created_at`
}
