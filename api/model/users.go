package model

import "time"

type User struct {
	ID        ID
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
